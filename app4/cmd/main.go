package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var counter atomic.Int64

func main() {
	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		count := counter.Add(1)

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "pong %d", count)
	})

	log.Println("Ping-pong server listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
