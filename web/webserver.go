package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	mux := mux.NewRouter()
	for _, route := range routes {
		mux.HandleFunc(route.route, route.handler).Methods(route.httpMethod)
	}

	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}

func serializeData(data interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if data, err := json.Marshal(data); err != nil {
		fmt.Fprintln(w, "{\"message\":\"Error reading!\"}")
		return err
	} else {
		fmt.Fprintln(w, string(data))
		return nil
	}
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}
