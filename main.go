package main

import (
	"fmt"
	"log"
	"net/http"
	"shobak/db"
	"shobak/pkg/setting"
	"shobak/routes"

	"github.com/kr/pretty"
)

func init() {
	setting.Setup("./config/config.json")
	db.Setup()
}

func main() {
	defer deferFunc()

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

func deferFunc() {
	pretty.Logln("[MAIN] Work has stopped!")
	db.CloseDB()
}
