package web

import (
	"net/http"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"goworkshop/persistence"
	"strconv"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	 WriteJson(w, GetAuthorsFromDB())
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	WriteJson(w, GetAuthorByIDFromDB(id))
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	DeleteAuthorByIDFromDB(id)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to create author: %s", err)
	} else {
		AddAuthorToDB(author)
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

func GetAuthorByIDFromDB(id int) model.Author {
	var author model.Author
	persistence.Connection.Where(&model.Author{Entity: model.Entity{ID: id}}).Find(&author)
	return author
}

func AddAuthorToDB(author model.Author) {
	err := persistence.Connection.Create(&author).Error
	if err != nil {
		fmt.Println(err)
	}
	return author.ID
}

func DeleteAuthorByIDFromDB(id int) {
	persistence.Connection.Delete(&model.Author{Entity: model.Entity{ID: id}})
}