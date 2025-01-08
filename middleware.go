package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Define the URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close() // Close the response body when done

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print the response
	fmt.Printf("Response: %s\n", body)

	mux := http.NewServeMux()
	helloHandler := http.HandlerFunc(HandleHello)
	mux.Handle("/hello", Middleware(helloHandler))

	http.ListenAndServe(":8080", mux)
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware: before")
		next.ServeHTTP(w, r)
		log.Println("Middleware: after")
	})
}
