package main

import (
	"flag"
	"./dropbox"
	"./google"
	"./yandex"
)

func main(){

	servicePtr := flag.String("service", "", "Text to parse.")
	loginPtr := flag.String("login", "", "Text to parse.")
	passwordPtr := flag.String("password", "", "")
	urlPtr := flag.String("url", "", "")
	flag.Parse()


	switch *servicePtr {
		case "dropbox":
			dropbox.Download(*loginPtr, *passwordPtr, *urlPtr)
		case "yandex":
			yandex.Download(*loginPtr, *passwordPtr, *urlPtr)
		case "google":
			google.Download(*loginPtr, *passwordPtr, *urlPtr)
	}
}
