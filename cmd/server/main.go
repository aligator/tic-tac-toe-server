package main

import (
	"github.com/aligator/tic-tac-toe-server/server"
	"log"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}
