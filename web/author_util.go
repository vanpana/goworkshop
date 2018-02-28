package web

import (
	"net/http"
	"fmt"
	"goworkshop/model"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	if err := serializeData(model.Authors, w); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading body!\"}")
		return
	}
	var author model.AuthorDto
	if err := json.Unmarshal(body, &author); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error unmarshling the body!\"}")
		return
	}
	model.Authors[author.UUID] = author
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	w.Header().Set("Content-Type", "application/json")

	author := model.Authors[uuid]

	if author != (model.AuthorDto{}) {
		if data, err := json.Marshal(author); err != nil {
			fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
			return
		} else {
			fmt.Fprintln(w, string(data))
			return
		}
	} else {
		fmt.Fprintln(w, "{\"message\":\"No author with this uuid!\"}")
		return
	}
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	w.Header().Set("Content-Type", "application/json")

	delete(model.Authors, uuid)

	return
}