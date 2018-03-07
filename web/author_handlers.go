package web

import (
	"net/http"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/persistence"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	 WriteJson(w, GetAuthorsFromDB())
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	authorUUID := mux.Vars(r)["uuid"]
	author, err := model.Authors.Get(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		WriteJson(w, author)
	}
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	var authorUUID = mux.Vars(r)["uuid"]
	err := model.Authors.Delete(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Failed to delete author: %s", err)
	} else {
		WriteJson(w, model.Authors)
	}
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to create author: %s", err)
	} else {
		model.Authors.Add(author)
		WriteJson(w, author)
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	author, err = model.Authors.Update(author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	WriteJson(w, author)
}

func GetAuthorsFromDB() model.AuthorsList{
	var authors model.AuthorsList
	persistence.Connection.Where(&model.Author{}).Find(&authors)
	return authors
}