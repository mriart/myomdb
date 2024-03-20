package main

import (
	"fmt"
	"net/http"

	"github.com/mriart/myomdb/movie"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("GET /movie", hMovie)

	fmt.Println("Server running at localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println((err))
	}
}

func hMovie(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query()["title"][0]
	m := movie.Movie{}
	err := m.GetMovie(t)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	str := m.PrintStringMovie()
	fmt.Fprintln(w, str)
}
