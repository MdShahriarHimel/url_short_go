package main

import (
	"log"
	"net/http"
	"url_short/config"
	"url_short/handlers/link"
	"url_short/handlers/user"
	"url_short/middleware"
)

func main() {
	mux := http.NewServeMux()

	config := config.GetConfig()

	manager := middleware.NewManager()

	manager.GlobalMiddlewareAssign(middleware.ApplicationJson)

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(mux, *manager)

	linkHandler := link.NewHandler()
	linkHandler.RegisterRoutes(mux, *manager)

	log.Printf("Creating Server at port :%s", config.PORT)

	portStr := ":" + config.PORT
	err := http.ListenAndServe(portStr, manager.GlobalMiddleWareWrapper(mux))

	if err != nil {
		log.Fatal("ERROR CREATING SERVER. EXITING...")
	}
}
