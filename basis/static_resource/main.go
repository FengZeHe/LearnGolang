package main

import "net/http"

func main() {
	serve := http.FileServer(http.Dir("."))
	http.ListenAndServe(":8080", serve)

}
