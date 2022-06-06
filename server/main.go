package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sharifsheraz/go-orders/db"
	"github.com/sharifsheraz/go-orders/handlers"
)

const ApiPort = 4000

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/orders", h.GetOrders).Methods(http.MethodGet)

	log.Println("API is running on port", ApiPort)
	http.ListenAndServe(fmt.Sprintf(":%d", ApiPort), router)
}
