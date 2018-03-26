package dropbox

import "fmt"

func Download(login string, password string, url string){
	println("This Dropbox download utility")
	fmt.Printf("login: %s, password: %s, url: %s\n", login, password, url)
}
