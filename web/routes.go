package web

import "net/http"

type Route struct {
	route      string
	handler    func(http.ResponseWriter, *http.Request)
	httpMethod string
}

var routes = []Route{
	// Books
	{
		route:      "/books",
		handler:    GetAllBooks,
		httpMethod: "GET",
	},
	{
		route:      "/books/{uuid}",
		handler:    GetBookByUUID,
		httpMethod: "GET",
	},
	{
		route:      "/books/{uuid}",
		handler:    UpdateBook,
		httpMethod: "PUT",
	},
	{
		route:      "/books/{uuid}",
		handler:    DeleteBookByUUID,
		httpMethod: "DELETE",
	},
	{
		route:      "/books",
		handler:    AddBook,
		httpMethod: "POST",
	},

	// Authors
	{
		route: "/authors/{uuid}",
		handler: getAuthor,
		httpMethod: "GET",
	},
	{
		route: "/authors",
		handler: createAuthor,
		httpMethod: "POST",
	},
	{
		route: "/authors",
		handler: getAllAuthors,
		httpMethod: "GET",
	},
	{
		route: "/authors",
		handler: updateAuthor,
		httpMethod: "PUT",
	},
	{
		route: "/authors/{uuid}",
		handler: deleteAuthor,
		httpMethod: "DELETE",
	},
}