package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintln(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Todo App</title>
			</head>
			<body>
				<h1>Hello from Kubernetes!</h1>
				<p>My first Go application running in Kubernetes.</p>
			</body>
			</html>
		`)
	})

	fmt.Printf("Server started in port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
