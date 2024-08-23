package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("$PORT must be set")
	}

	router := chi.NewRouter()

	// Cores setup
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, http.StatusBadRequest, "This is an error")
	})

	router.Mount("/v1", v1Router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	fmt.Printf("Server is running on port: http://localhost:%v\n", PORT)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
