package dropbox

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const PREFIX = "https://www.dropbox.com/s/"
const POSTFIX = "?dl=1"

func Download(login string, password string, url string) (err error) {
	fmt.Print("This Dropbox download utility\n")
	//fmt.Printf("login: %s, password: %s, url: %s\n", login, password, url)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter fileID: ")
	fileID, _ := reader.ReadString('\n')
	fmt.Print("Enter file name: ")
	fileName, _ := reader.ReadString('\n')

	s := []string{PREFIX,
		strings.TrimSuffix(fileID, "\n"),
		"/",
		strings.TrimSuffix(fileName, "\n"),
		POSTFIX}

	URI := strings.Join(s, "")

	return Run(URI, strings.TrimSuffix(fileName, "\n"))
}

func Run(url string, fileName string) (err error) {
	fmt.Printf("RUN got a new URL request: %s\n", url)
	// check path downloads exists
	if _, err := os.Stat("downloads"); os.IsNotExist(err) {
		os.MkdirAll("downloads", os.ModePerm)
	}

	var buffer bytes.Buffer
	buffer.WriteString("downloads/")
	buffer.WriteString(fileName)

	out, err := os.Create(buffer.String())
	if err != nil {
		return err
	}
	defer out.Close()

	println("file downloading....")
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
	if err != nil {
		return err
	}
	println("Success: file saved on disk!")
	return nil
}
