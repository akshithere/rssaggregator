package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
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
		v1Router.Get("/health", handlerReadiness)
		v1Router.Get("/err", handlerError)

		router.Mount("/v1", v1Router)

		log.Printf("Server is running healthily")
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("Somethings up", err)
		}
	http.ListenAndServe(port,nil)
}
