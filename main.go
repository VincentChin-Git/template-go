package main

import (
	"client-server-go/middleware"
	"client-server-go/routes"
	"client-server-go/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func main() {

	fmt.Println("Initing Router...")
	r := chi.NewRouter()

	fmt.Println("Setting Up Middleware...")
	r = middleware.SetupMiddleware(r) // setup middleware

	fmt.Println("Setting Up Routers...")
	r.Mount("/api", routes.MainRouter()) // setup router

	fmt.Println("Connecting To Redis...")
	storage.ConnectRedis()

	fmt.Println("Connecting To MongoDB...")
	storage.ConnectDatabase()

	fmt.Println("Server Started!") // todo: add this log after the server starts

	err := http.ListenAndServe(":5050", r)

	if err != nil {
		log.Fatalf(":( Server Start Error: %s", err)
	}

}
