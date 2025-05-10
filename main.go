package main

import (
    "log"
    "net/http"
)

type homeHandler struct {
}

func (hh homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome!"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is the about page."))
} 

func main() {
    mux := http.NewServeMux()

    hh := homeHandler{}

	// hh does implement ServeHTTP, just pass it to mux
    mux.Handle("/home", hh)

	// because aboutHandler does not implement ServeHTTP so we pass it to HandlerFunc cause HandlerFunc does implement ServeHTTP. So aboutHandler is now a handlerFunc
    mux.HandleFunc("/about", aboutHandler)

    log.Fatal(http.ListenAndServe(":8080", mux))
}    