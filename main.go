package main

import (
	"log"

	"github.com/AYDEV-FR/ISEN-Api/pkg/api"
)

func main() {
	log.Printf("Server started")

	router := api.NewRouter()

	log.Fatal(router.Run(":8080"))
}
