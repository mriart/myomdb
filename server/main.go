package main

import (
	"fmt"
	"net/http"
)

func main() {
	hFileServer := http.FileServer(http.Dir("."))
	http.Handle("/", hFileServer)

	fmt.Println("Server running at localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println((err))
	}
}
