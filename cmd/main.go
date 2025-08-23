package main

import (
	"extrator/internal/config"
	"log"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v", config)
}
