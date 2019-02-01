package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
	fmt.Println(r.URL)
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func about(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
	fmt.Println(r.URL.Path[1:])
    fmt.Fprintf(w, "<h1>Welcome to the about page!</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
    log.Fatal(http.ListenAndServe(":8080", nil))
}



func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}