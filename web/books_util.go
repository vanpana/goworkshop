package web

import (
	"net/http"
	"github.com/gorilla/mux"
	"goworkshop/model"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, model.Books)
}

func GetBookByUUID(w http.ResponseWriter, r *http.Request) {
	var bookUUID = mux.Vars(r)["uuid"]
	book := model.Books[bookUUID]
		WriteJson(w, book)
}

func DeleteBookByUUID(w http.ResponseWriter, r *http.Request) {
	var bookUUID = mux.Vars(r)["uuid"]
	delete(model.Books, bookUUID)
	WriteJson(w, model.Books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.BookDto
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to create book: %s", err)
	} else {
		model.Books[book.UUID] = book
		WriteJson(w, book)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book model.BookDto
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to update book: %s", err)
		return
	}
	model.Books[book.UUID] = book
	if err != nil {
		fmt.Fprintf(w, "Failed to update book: %s", err)
		return
	}
	WriteJson(w, book)
}