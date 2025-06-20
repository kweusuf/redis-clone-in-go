package main

import (
	"log/slog"

	"github.com/kweusuf/redis-clone-in-go/boot"
)

// main is the entry point of the application.
// It logs a startup message and initializes the Redis clone server.
func main() {
	slog.Info("Hello World! Starting Redis...")
	boot.Init()
}
