package main

import (
	"os"
	"net/http"
	"io"
	"fmt"
)

func main(){
	//url := "https://drive.google.com/open?id=161GOPMpvTR3ke3_PmQLK3YITa_99TZ9u"
	url := "https://edaplus.info/food_pictures/bream.jpg"
	filepath:="C:/temp"
	err := downloadFile(filepath, url)
	if err != nil  {
		println(err)
	}
}
func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create("C:/temp/tempus.txt")
	if err != nil  {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		return err
	}

	return nil
}
