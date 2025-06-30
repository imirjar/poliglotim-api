package main

import (
	"context"
	"log"

	"github.com/imirjar/poliglotim-api/internal/app"
)

func main() {
	ctx := context.Background()
	log.Fatal(app.Start(ctx))
}
