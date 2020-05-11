package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "github.com/jaydeep87/poc-go-pg/config"
	routes "github.com/jaydeep87/poc-go-pg/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8081"))
}
