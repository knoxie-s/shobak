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

	fmt.Println("1")
	log.Println("1")

	// service connections
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen:", err)
	}

}
