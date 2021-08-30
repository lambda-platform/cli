package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	humanize "github.com/dustin/go-humanize"
	//"time"
)

func main() {
	var showVersion bool
	var createApp bool
//
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.BoolVar(&createApp, "create", false, "Create Lambda project.")
	flag.BoolVar(&createApp, "c", false, "Create Lambda project.")
	flag.Parse()

	// Show version and exit
	if showVersion {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	}
	if(createApp){
		err := DownloadGStarter()

		if(err != nil){
			fmt.Println(err)
		} else {
			fmt.Println("\nDownload Completed")
		}
		errUnZip := UnZipLambdaCodes()
		if(errUnZip != nil){
			fmt.Println(errUnZip)
		}
	}
}

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
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
func DownloadGStarter() error{
	url := "https://lambda.cloud.mn/starter.zip"
	fmt.Println("Download Started")
	// Create the file
	out, err := os.Create("lambda.zip")
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


func UnZipLambdaCodes() error{

	_, fileName, _, _ := runtime.Caller(0)


	var dest string = filepath.Dir(fileName)

	fmt.Println(dest)

	var src string = "lambda.zip"


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
		return nil







}
