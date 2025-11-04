package main

import (
	"log"
	"net/http"
	"os"

	_ "ordersystem/docs" // keeps import happy even if docs is empty
	"ordersystem/repository"
	"ordersystem/rest"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // loads .env if present (harmless in Docker)

	store, err := repository.NewStore()
	if err != nil {
		log.Fatalf("init store: %v", err)
	}

	appPort := getenv("APP_PORT", "3000")
	srv := rest.NewServer(store)
	log.Printf("Order System starting on :%s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, srv))
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
