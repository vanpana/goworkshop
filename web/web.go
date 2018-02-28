package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"goworkshop/model"
	"github.com/gorilla/mux"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{uuid}", getBookByUUID).Methods("GET")
	router.HandleFunc("/books/{uuid}/author", getBookAuthor).Methods("GET")

	router.HandleFunc("/author", createAuthor).Methods("POST")


	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := json.Marshal(model.Books)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(books))
}
func getAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := json.Marshal(model.Authors)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(authors))
}

func getBookByUUID(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	for _, book := range model.Books {
		if book.UUID == uuid {
			data, err := json.Marshal(book)

			if err != nil {
				panic(err)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintln(w, string(data))
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Book not found")
}
func getBookAuthor(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	for _, book := range model.Books {
		if book.UUID == uuid {
			data, err := json.Marshal(book.Author)

			if err != nil {
				panic(err)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintln(w, string(data))
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Book not found")
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	r.Body.Read()
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}