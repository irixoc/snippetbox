package main

import (
	"log"
	"net/http"
)

// home handler
// http.ResponseWriter parameter provides methods for assembling a HTTP response and sending it to the user
// *http.Request parameter is a pointer to a struct which holds information about the current request
// (like the HTTP method and the URL being requested)
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	// This ensures paths with trailing slashes (/contact/) don't get directed to the home page
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello, World"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("showing you a snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating a snippet"))
}

func main() {
	// http.NewServeMux() function initializes a new servemux
	mux := http.NewServeMux()

	// register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", showSnippet)

	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass in two parameters:
	// the TCP network address to listen on (in this case ":4000"),
	// and the servemux we just created.
	err := http.ListenAndServe(":4000", mux)

	// If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Fatal(err)
}
