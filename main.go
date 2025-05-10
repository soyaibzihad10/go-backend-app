package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	/*
		/books/all এবং /books/{isbn} এই দুইটি route একে অপরের সঙ্গে conflict করতে পারে path matching এর সময়।
		এখন, ইউজার যদি ব্রাউজারে /books/all হিট করে...

		👉 রাউটার কীভাবে জানবে যে এটা all নামে একটি ফিক্সড রাউট, নাকি এটি {isbn} নামে কোনো স্ট্রিং ভ্যালু?
		
		Soulution: Go-এর gorilla/mux রাউটার উপরের নিয়মে কাজ করে:
		It matches in order from top to bottom.
	*/

	booksSSubR := r.PathPrefix("/books").Subrouter()
	booksSubR.HandleFunc("/all", AllHandler).Methods(http.MethodGet)
    booksSubR.HandleFunc("/{isbn}", IspnHandler).Methods(http.MethodGet)
    booksSubR.HandleFunc("/new", NewHandler).Methods(http.MethodPost)
    booksSubR.HandleFunc("/update", UpdateHandler).Methods(http.MethodPut)
    booksSubR.HandleFunc("/delete/{isbn}", DeleteIspnHandler).Methods(http.MethodDelete)

    log.Fatal(http.ListenAndServe(":8090", r))
}
