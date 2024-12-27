package main

import (
	"fmt"
	"log"

	"github.com/bupd/digital-wellbeing/pkg/server"
)

func main() {
	server := server.NewServer()

	fmt.Printf("Ground Control running on port %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
