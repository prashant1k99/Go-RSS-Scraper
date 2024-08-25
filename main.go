package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prashant1k99/Go-RSS-Scraper/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("$PORT must be set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("$DB_URL must be set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	apiConfig := apiConfig{
		DB: database.New(conn),
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
	v1Router.Post("/user", apiConfig.createUser)
	v1Router.Get("/user", apiConfig.middlewareAuth(apiConfig.handleGetUser))
	v1Router.Post("/feed", apiConfig.middlewareAuth(apiConfig.createFeed))
	v1Router.Get("/feed", apiConfig.middlewareAuth(apiConfig.getFeedByUser))
	v1Router.Get("/feed/{feedId}", apiConfig.middlewareAuth(apiConfig.getFeedById))
	v1Router.Put("/feed/{feedId}", apiConfig.middlewareAuth(apiConfig.updateFeedById))
	v1Router.Delete("/feed/{feedId}", apiConfig.middlewareAuth(apiConfig.deleteFeedById))

	router.Mount("/v1", v1Router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	fmt.Printf("Server is running on port: http://localhost:%v\n", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
