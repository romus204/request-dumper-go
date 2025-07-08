package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func logRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----- New query -----")
	fmt.Printf("Mathod: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL.String())
	fmt.Printf("Path: %s\n", r.URL.Path)
	fmt.Printf("Query params: %s\n", r.URL.RawQuery)

	ip := r.RemoteAddr
	fmt.Printf("IP client: %s\n", ip)

	fmt.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", name, value)
		}
	}

	if r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error read body: %v\n", err)
		} else {
			body := strings.TrimSpace(string(bodyBytes))
			if len(body) > 0 {
				fmt.Println("Body:")
				fmt.Println(body)
			} else {
				fmt.Println("Body: empty")
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", logRequest)

	port := "8085"
	fmt.Printf("Server start on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
