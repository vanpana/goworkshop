package main

import (
	"fmt"
	"io/ioutil"
	"goworkshop/model"
	"encoding/json"
	"goworkshop/web"
)

func main() {
	booksFileData, booksErr := ioutil.ReadFile("data/books.json")
	authorsFileData, authorsErr := ioutil.ReadFile("data/authors.json")

	if booksErr != nil {
		panic(booksErr)
	}

	if authorsErr != nil {
		panic(authorsErr)
	}

	if booksUnmarshalErr := json.Unmarshal(booksFileData, &model.Books); booksUnmarshalErr != nil {
		panic(booksUnmarshalErr)
	}

	if authorsUnmarshalErr := json.Unmarshal(authorsFileData, &model.Authors); authorsUnmarshalErr != nil {
		panic(authorsUnmarshalErr)
	}


	fmt.Println(model.Books)
	fmt.Println(model.Authors)

	web.StartServer()



}
