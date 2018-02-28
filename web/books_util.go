package web

import (
	"net/http"
	"github.com/gorilla/mux"
	"goworkshop/model"
	"fmt"
	"encoding/json"
)

func getBookAuthor(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]
	for _, book := range model.Books {
		if book.UUID == uuid {
			if err := serializeData(book.Author, w); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}

	fmt.Fprintln(w, "{\"message\":\"The book does not exist!\"}")
	w.WriteHeader(http.StatusNotFound)
}

func getBookByUuid(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	w.Header().Set("Content-Type", "application/json")

	for _, book := range model.Books {
		if book.UUID == uuid {
			if data, err := json.Marshal(book); err != nil {
				fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
				return
			} else {
				fmt.Fprintln(w, string(data))
				return
			}
		}
	}

	fmt.Fprintln(w, "{\"message\":\"The book does not exist!\"}")
	w.WriteHeader(http.StatusNotFound)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	w.Header().Set("Content-Type", "application/json")

	if data, err := json.Marshal(model.Books); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
	} else {
		fmt.Fprintln(w, string(data))
	}

}