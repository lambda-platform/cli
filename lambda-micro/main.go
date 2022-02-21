package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	humanize "github.com/dustin/go-humanize"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	//"time"
)

func main() {
	var showVersion bool
	var showHelp bool
	var key string

	//
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.BoolVar(&showHelp, "help", false, "Print version information.")
	flag.BoolVar(&showHelp, "h", false, "Print version information.")
	flag.StringVar(&key, "k", "", "Lambda project key.")
	flag.StringVar(&key, "key", "", "Lambda project key.")

	flag.Parse()

	// Show version and exit
	if showVersion {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	} else if showHelp {
		ShowHelp()
	} else {
		args := flag.Args()

		if(len(args) >= 1){
			var command string = flag.Args()[0]



			if(command == "c" || command == "create"){
				if(len(args) >= 2){
					var projectName string = flag.Args()[1]
					dir, _ := os.Getwd()
					dest := dir+"/"+projectName
					if _, err := os.Stat(dest); os.IsNotExist(err) {
						fmt.Println("Creating: "+projectName)
						os.MkdirAll(dest, 0755)

						err := DownloadGStarter(dest)

						if(err != nil){
							fmt.Println(err)
						} else {
							fmt.Println("\nDownload Completed")
						}
						errUnZip := UnZipLambdaCodes(dest)
						if(errUnZip != nil){
							fmt.Println(errUnZip)
						}

						if(key != ""){
							UpdateProjectKey(dest, key)
						}

					} else {
						fmt.Println("Project already exist in: "+dir)
					}

				} else {
					fmt.Println("Please insert project name")
				}

			} else {
				ShowHelp()
				fmt.Println("Unknown command: "+command)
			}


		}




	}

}

func UpdateProjectKey(dest string, key string)  {
	lambdaConfig, _ := ioutil.ReadFile(dest + "/lambda.json")
	lambdaConfigContent := strings.ReplaceAll(string(lambdaConfig), "\"project_key\": \"\"", "\"project_key\": \""+key+"\"")
	WriteFile(lambdaConfigContent, dest + "/lambda.json")
}
func WriteFile(fileContent string, path string) {

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	l2, err := f.WriteString(string(fileContent))
	if err != nil {
		panic(err)
		panic(l2)
		f.Close()
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

type WriteCounter struct {
	Total uint64
}

func ShowHelp()  {
	help:= fmt.Sprintf(`Usage: lambda <comman> [options]

Options:
-v, --version                      output the version number
-h, --help                         output usage information
-k, --key                          project key

Commands:
[options] create <app-name>        create a new project
[options] c <app-name>             create a new project

Examples:
lambda create my-app
lambda -key=<project-key> create my-app 
`)
	fmt.Println(help)
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}
func DownloadGStarter(dest string) error{
	url := "https://lambda.cloud.mn/microservice.zip"
	fmt.Println("Download Started")
	// Create the file
	out, err := os.Create(dest+"/lambda.zip")
	if err != nil {

		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}


	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}


	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("file not found error")
	}


	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)




	return err
}
func UnZipLambdaCodes(dest string) error{


	var src string = dest+"/"+"lambda.zip"


	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return  err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return  err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return  err
		}
	}

	e := os.Remove(src)
	if e != nil {
		return e
	}


	return nil


}
