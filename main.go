package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = ":8080"
	}
	fmt.Printf("Listening on port: %v", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from go and docker!")
	})
	fmt.Println(http.ListenAndServe(port, nil))
}
