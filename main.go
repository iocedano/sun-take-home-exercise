package main

import (
	API "interview/takehomeproject/server/v2"
	"net/http"
	"path"
)

var sources = map[string]string{
	"google": "https://www.google.com",
	"amazon": "https://www.amazon.com",
}

func main() {
	mux := http.NewServeMux()

	// Client
	buildPath := path.Clean("client/build")
	mux.Handle("/", http.FileServer(http.Dir(buildPath)))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("client/build/static"))))

	// Rest API V1
	API.SetupRouters(mux, sources)

	http.ListenAndServe(":8080", mux)
}
