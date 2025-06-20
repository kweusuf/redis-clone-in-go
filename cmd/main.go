package main

import (
	"log/slog"

	"github.com/kweusuf/redis-clone-in-go/boot"
)

func main() {
	slog.Info("Hello World! Starting Redis...")
	boot.Init()
}
