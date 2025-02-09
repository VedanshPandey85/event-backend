package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-event-planner/handlers"
	"weather-event-planner/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.InitDB()
}

func main() {
	r := mux.NewRouter()

	// Event Management Service
	r.HandleFunc("/api/events", handlers.CreateEvent).Methods("POST")
	r.HandleFunc("/api/events", handlers.GetEvents).Methods("GET")
	r.HandleFunc("/api/events/{id}", handlers.GetEvent).Methods("GET")
	r.HandleFunc("/api/events/{id}", handlers.UpdateEvent).Methods("PUT")
	r.HandleFunc("/api/events/{id}", handlers.DeleteEvent).Methods("DELETE")

	// Weather Fetch Service
	r.HandleFunc("/api/weather", handlers.FetchWeather).Methods("GET")

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
