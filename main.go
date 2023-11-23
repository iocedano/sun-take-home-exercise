package main

import (
	v1 "interview/takehomeproject/server/v1"
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
	v1.SetupRouters(mux, sources)

	http.ListenAndServe(":8080", mux)
}
