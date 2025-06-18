package main

import (
	"fmt"
	"log"

	"github.com/imirjar/poliglotim-api/config"
	"github.com/imirjar/poliglotim-api/internal/app"
)

func main() {
	config := config.New()
	fmt.Printf("Starting server on the port %s... \n", config.Port)
	log.Fatal(app.Start())
}
