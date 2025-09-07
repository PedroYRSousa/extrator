package main

import (
	"extrator/internal/configs"
	"log"
)

func main() {
	log.Println("Teste")
	_, err := configs.Load()
	if err != nil {
		log.Println(err)
	}
}
