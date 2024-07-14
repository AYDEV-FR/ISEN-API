package main

import (
	"log"

	_ "github.com/AYDEV-FR/ISEN-Api/docs"
	"github.com/AYDEV-FR/ISEN-Api/pkg/api"
)

//	@title			ISEN-API
//	@version		0.0.5
//	@description	REST API scrapping the webaurion website from ISEN Méditerranée

//	@contact.name	GitHub
//	@contact.url	https://github.com/AYDEV-FR/ISEN-API

//	@license.name	MIT License
//	@license.url	https://github.com/AYDEV-FR/ISEN-API/blob/main/LICENSE

//	@BasePath	/v1

// @securitydefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Token
// @description				Token from ISEN webaurion to access all the user's data
func main() {
	log.Printf("Server started")

	router := api.NewRouter()

	log.Fatal(router.Run(":8080"))
}
