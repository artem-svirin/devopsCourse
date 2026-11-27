package main

import (
	"fmt"
	"time"
	"uuid"
)

func main() {
	uuid := uuid.New()
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	for {
		// logger.Info(uuid.String())
		fmt.Printf("%s: %s\n", time.Now().UTC().Format(time.RFC3339Nano), uuid)
		time.Sleep(5 * time.Second)
	}
}
