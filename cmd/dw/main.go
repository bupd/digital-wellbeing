package main

import (
	"fmt"
	"log"

	"github.com/bupd/digital-wellbeing/pkg/server"
  _ "github.com/joho/godotenv/autoload"
)

func main() {
	server := server.NewServer()

  fmt.Printf("\nDigital Wellbeing running on server: %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
