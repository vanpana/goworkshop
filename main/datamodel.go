package main

import "fmt"

//BookDto - The DTO used to access books
type BookDto struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	NoPages     int `json:"no_pages"`
	ReleaseDate string `json:"release_date"`
	Author      AuthorDto `json:"author"`
}

func (b BookDto) String() string {
	return fmt.Sprintf("BookDto(UUID=%s, Title=%s, Nopages=%d, ReleaseDate=%s, Author=$s)",
		b.UUID, b.Title, b.NoPages, b.ReleaseDate, b.Author)
}

//AuthorDto - The DTO used to access authors
type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

func (a AuthorDto) String() string {
	return fmt.Sprintf("AuthorDto(UUID=%s, First Name=%s, Last Name=%s, BirthDay=%s, Death=$s)",
		a.UUID, a.FirstName, a.LastName, a.Birthday, a.Death)
}


//Books - the list of available books
var Books = []BookDto{}

// Authors - the list of available authors
var Authors = []AuthorDto{}
