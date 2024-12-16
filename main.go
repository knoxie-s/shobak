package main

import (
	"fmt"
	"log"
	"net/http"
	"shobak/pkg/setting"
	"shobak/routes"
)

func init() {
	setting.Setup("./config/config.json")
}

func main() {
	endPoint := fmt.Sprintf(":%d", setting.Config.App.Port)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routes.Init(),
	}

	// service connections
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen:", err)
	}

}
