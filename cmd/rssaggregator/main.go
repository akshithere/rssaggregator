package main

import (
	"log"
	"net/http"
	"os"

	// "database/sql"
	"github.com/akshithere/rssaggregator/internal/db"
	"github.com/akshithere/rssaggregator/pkg/config"
	"github.com/akshithere/rssaggregator/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load environment variables from .env file
	port := os.Getenv("PORT") // Get the PORT from environment variables

	db.ConnectToDB()
	config.InitConfig()
	
	if port == ""{
		log.Fatal("Port is not provided")
	}
		// creates a router
		router:= chi.NewRouter()
		srv := &http.Server{
			Handler: router,
			Addr: ":"+port,
		}

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string {"https://*","http://*"},
			AllowedMethods: []string {"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string {"*"},
			ExposedHeaders: []string {"Link"},
			AllowCredentials: false,
			MaxAge: 300,
		}))

		v1Router := chi.NewRouter()
		v1Router.Get("/health", handlers.HandlerReadiness)
		v1Router.Get("/err", handlers.HandlerError)
		v1Router.Post("/createUserTable", handlers.HandlerCreateUserTable)

		router.Mount("/v1", v1Router)

		log.Printf("Server is running healthily")
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("Somethings up", err)
		}
	http.ListenAndServe(port,nil)
		defer config.ApiConfig.Database.Close()
}
