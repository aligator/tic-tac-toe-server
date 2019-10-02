package server

import "log"

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}
