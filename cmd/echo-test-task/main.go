package main

import (
	"log"

	"github.com/ArtificialAI/echo-handler/internal/pkg/app"
)

func main() {
	// Echo instance
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}

}
