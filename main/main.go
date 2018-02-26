package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	bookFileContent, err := ioutil.ReadFile("main/books.json")

	if err != nil {
		//fmt.Println("Error occured")
		//os.Exit(1)

		panic(err) // Similar to the lines above
	}

	var books []BookDto

	if err = json.Unmarshal(bookFileContent, &books); err != nil {
		panic(err)
	}

	fmt.Println(books)

	serializedData, err := json.Marshal(books)

	if err != nil {
		panic(err)
	}

	fmt.Println("Serialized Books are: ")
	fmt.Println(string(serializedData))


}
