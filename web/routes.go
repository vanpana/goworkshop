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
		handler:    getAllBooks,
		httpMethod: "GET",
	},
	{
		route:      "/books/{uuid}",
		handler:    getBookByUuid,
		httpMethod: "GET",
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
}