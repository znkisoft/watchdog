package main

import (
	"context"
	"log"

	"github.com/znkisoft/watchdog/server"
)

func main() {
	s := server.Server{}
	ctx := context.Background()
	if err := s.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
