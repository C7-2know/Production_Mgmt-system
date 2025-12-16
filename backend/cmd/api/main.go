package main

import (
	"log"
	"fmt"

	"soap_factory/cmd/api/bootstrap"
	"soap_factory/delivery/http"
)

func main() {
	app := bootstrap.App()
	defer app.Close()

	log.Printf("Starting server on port %d...", app.Env.Port)
	
	router := http.SetupRoutes(app.DB)

	err := router.Run(fmt.Sprintf(":%d", app.Env.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
