package dropbox

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"bufio"
	"bytes"
	"strings"
)

const	PREFIX = "https://www.dropbox.com/s/"
const 	POSTFIX = "?dl=1"

func Download(login string, password string, url string) (err error){
	println("This Dropbox download utility")
	fmt.Printf("login: %s, password: %s, url: %s\n", login, password, url)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	URI, _ := reader.ReadString('\n')
	fmt.Print("Enter file name: ")
	fileName, _ := reader.ReadString('\n')

	fmt.Println(URI)
	//s := []string{"https://www.dropbox.com/s/", strings.TrimSuffix(fileID, "\n"), "/", strings.TrimSuffix(fileName, "\n"), "?dl=1"}
	//s := []string{"https://www.dropbox.com/s/", "1uohkqwl2k2d1du/anyconnect.rar", "?dl=1"}
	//s := []string{"https://www.dropbox.com/s/", strings.TrimSuffix(fileID, "\n"), "?dl=1"}
	/*s := []string{"this", strings.TrimSuffix(fileID, "\n"), "is", "joined", "string\n"}
		URI := strings.Join(s, "")*/

	//Run(URI, strings.TrimSuffix(fileName, "\n"))
	return Run("https://www.dropbox.com/s/1uohkqwl2k2d1du/anyconnect.rar?dl=1", strings.TrimSuffix(fileName, "\n"))
	//return nil
}

func Run(url string, fileName string) (err error){
	fmt.Printf("RUN got a new URL request: %s\n", url)

	var buffer bytes.Buffer
	buffer.WriteString("downloads/")
	buffer.WriteString(fileName)

	out, err := os.Create(buffer.String())
	if err != nil  {
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
	if err != nil  {
		return err
	}
	println("Success: file saved on disk!")
	return nil
}
