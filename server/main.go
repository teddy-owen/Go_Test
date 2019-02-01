package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[1:])
    fmt.Fprintf(w, "<h1>Welcome to the about page!</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
    log.Fatal(http.ListenAndServe(":8080", nil))
}