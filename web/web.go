package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"goworkshop/model"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/authors", getAuthors)
	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
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
func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}