package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"io/ioutil"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

type Route struct {
	route string,
	handler func(http.ResponseWriter, *http Request)
	httpMethod string
}

var routes = []Route {
	{
		route: "/books",
		handler: getBooks,
		httpMethod: "GET",

	},
}
func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{uuid}", getBookByUUID).Methods("GET")
	router.HandleFunc("/books/{uuid}/author", getBookAuthor).Methods("GET")

	router.HandleFunc("/author", createAuthor).Methods("POST")
	router.HandleFunc("/author", getAuthors).Methods("GET")


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
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading body\"")
		return
	}

	var author model.AuthorDto

	// TODO: fix empty author
	if err = json.Unmarshal(body, &author); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error unmarshaling the body\"")
		return
	}

	model.Authors = append(model.Authors, author)
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}