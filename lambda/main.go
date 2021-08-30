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
	"strings"
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
		DownloadGStarter()
		UnZipLambdaCodes()
	}
}

func DownloadGStarter() error{
	url := "https://lambda.cloud.mn/starter.zip"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}


	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("file not found error")
	}

	// Create the file
	out, err := os.Create("lambda.zip")
	if err != nil {

		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)


	return err
}


func UnZipLambdaCodes() error{
	var dest string = ""
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