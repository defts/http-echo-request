package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func echoMyRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		fmt.Printf("%s: %s\n", name, headers)
	}
	// If this is a POST, add post data
	if r.Method == http.MethodPost {
		fmt.Printf("\n")
		r.ParseForm()
		fmt.Printf("%s\n", r.Form)
	}

	fmt.Printf("\n")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	http.HandleFunc("/", echoMyRequest)
	err := http.ListenAndServe(fmt.Sprintf(":%s", getenv("PORT", "9090")), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
