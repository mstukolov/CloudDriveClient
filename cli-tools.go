package main

import (
	"flag"
	"./download"
)

func main(){
	println("Started Dropbox tool")

	loginPtr := flag.String("login", "", "Text to parse.")
	passwordPtr := flag.String("password", "", "")
	urlPtr := flag.String("url", "", "")
	flag.Parse()

	//fmt.Printf("login: %s, password: %s, url: %s\n", *loginPtr, *passwordPtr, *urlPtr)
	download.Download(*loginPtr, *passwordPtr, *urlPtr)
}
