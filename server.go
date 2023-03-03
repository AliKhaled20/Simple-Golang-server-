package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If the request is for the root path, serve the index.html file
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}
		
		// If the request is for a path ending with .html, serve the corresponding file
		if r.URL.Path[len(r.URL.Path)-5:] == ".html" {
			http.ServeFile(w, r, r.URL.Path[1:])
			return
		}
		
		// If the request is for a path ending with .css, serve the corresponding file
		if r.URL.Path[len(r.URL.Path)-4:] == ".css" {
			http.ServeFile(w, r, r.URL.Path[1:])
			return
		}
		
		// If the request is for a path ending with .js, serve the corresponding file
		if r.URL.Path[len(r.URL.Path)-3:] == ".js" {
			http.ServeFile(w, r, r.URL.Path[1:])
			return
		}
		
		// If the request is for a path ending with .png, serve the corresponding file
		if r.URL.Path[len(r.URL.Path)-4:] == ".png" {
			http.ServeFile(w, r, r.URL.Path[1:])
			return
		}
		// If the request is for a path ending with .jpg, serve the corresponding file
		if r.URL.Path[len(r.URL.Path)-4:] == ".jpg" {
			http.ServeFile(w, r, r.URL.Path[1:])
			return
		}
		
		// Check the requested path and redirect to the appropriate website
		switch r.URL.Path {
		case "/gl":
			http.Redirect(w, r, "https://www.google.com/", http.StatusTemporaryRedirect)
			return
		case "/ghb":
			http.Redirect(w, r, "https://www.github.com/", http.StatusTemporaryRedirect)
			return
		case "/bzu":
			http.Redirect(w, r, "https://www.birzeit.edu/", http.StatusTemporaryRedirect)
			return
		}
		
		// Check if the file exists and serve it, or return an error page
		if _, err := os.Stat(r.URL.Path[1:]); err == nil {
			http.ServeFile(w, r, r.URL.Path[1:])
			return
		}

		// Set the response status to 404 Not Found
		w.WriteHeader(http.StatusNotFound)

		// Write the error page
		fmt.Fprintf(w, "<!DOCTYPE html>\n")
		fmt.Fprintf(w, "<html>\n")
		fmt.Fprintf(w, "<head>\n")
		fmt.Fprintf(w, "<title>Error</title>\n")
		fmt.Fprintf(w, "</head>\n")
		fmt.Fprintf(w, "<body>\n")
		fmt.Fprintf(w, "<p style='color: red'>The file is not found</p>\n")
		fmt.Fprintf(w, "<p><strong>Your names and IDs</strong></p>\n")
		fmt.Fprintf(w, "<p>The IP and port number of the client: %s</p>\n", r.RemoteAddr)
		fmt.Fprintf(w, "</body>\n")
		fmt.Fprintf(w, "</html>\n")


	// Otherwise, handle the request as before
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Listen on port 6677
	http.ListenAndServe(":6677", nil)
}
