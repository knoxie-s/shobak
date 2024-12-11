package main

import (
	"fmt"
	"log"
	"net/http"
	"shobak/routes"
)

func main() {
	endPoint := fmt.Sprintf(":%d", 8080)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routes.Init(),
	}

	// service connections
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen:", err)
	}

}
