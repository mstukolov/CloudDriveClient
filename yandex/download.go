package yandex

import "fmt"

func Download(login string, password string, url string){
	println("This Yandex Drive download utility")
	fmt.Printf("login: %s, password: %s, url: %s\n", login, password, url)
}
