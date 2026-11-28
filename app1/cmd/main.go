package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Status struct {
	Timestamp    string `json:"timestamp"`
	RandomString string `json:"randomString"`
}

func main() {
	// Генерируем UUID один раз при запуске.
	randomString := uuid.New().String()

	// Периодически выводим состояние в логи.
	go func() {
		for {
			fmt.Printf("%s: %s\n",
				time.Now().UTC().Format(time.RFC3339Nano),
				randomString,
			)

			time.Sleep(5 * time.Second)
		}
	}()

	// HTTP endpoint.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		status := Status{
			Timestamp:    time.Now().UTC().Format(time.RFC3339Nano),
			RandomString: randomString,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(status); err != nil {
			log.Printf("failed to encode response: %v", err)
		}
	})

	log.Println("HTTP server listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
