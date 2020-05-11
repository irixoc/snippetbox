package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// home handler
// http.ResponseWriter parameter provides methods for assembling a HTTP response and sending it to the user
// *http.Request parameter is a pointer to a struct which holds information about the current request
// (like the HTTP method and the URL being requested)
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// This ensures paths with trailing slashes (/contact/) don't get directed to the home page
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		// IMPORTANT: we must return from the handler. If we don't return, the handler
		// would keep executing and also write the "Hello from SnippetBox" message.
		return
	}

	w.Write([]byte("Hello, World"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Println(formatRequest(r))
	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Displaying snippet with id: %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	// Use r.Method to check whether the request is using POST or not.
	// Note: http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {

		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		// Note: call this method to change reponse header map before calling w.WriteHeader() and w.Write()
		w.Header().Set("Allowed", http.MethodPost)

		http.Error(w, "Method Not Allowed. POST requests only", 405)

		// Return from the function so that the subsequent code is not executed
		return
	}

	w.Write([]byte("creating a snippet"))
}

func main() {
	// http.NewServeMux() function initializes a new servemux (or router)
	mux := http.NewServeMux()

	// register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", showSnippet)

	mux.HandleFunc("/snippet/create/", createSnippet)

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

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
