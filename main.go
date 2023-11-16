package main

import (
	"fmt"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email us at <a href=\"mailto:info@zero-one-group.com\">info@zero-one-group.com</a>.</p>")

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
