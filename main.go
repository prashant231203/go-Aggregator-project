package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prashant231203/go-Aggregator-project/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	feed, err := urlToFeed("https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)
	// Load environment variables from .env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port from environment variables
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can't connect to databse")
	}

	db := database.New(conn)
	apicfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apicfg.handlerCreateUser)
	v1Router.Post("/feeds", apicfg.middlewareAuth(apicfg.handlerCreateFeed))
	v1Router.Get("/feeds", apicfg.handlerGetFeeds)
	v1Router.Post("/feed_follows", apicfg.middlewareAuth(apicfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apicfg.middlewareAuth(apicfg.handlerGetFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", apicfg.middlewareAuth(apicfg.handlerDeleteFeedFollow))
	v1Router.Get("/posts", apicfg.middlewareAuth(apicfg.handlerGetPostsForUser))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running on port", portString)
	http.ListenAndServe(":"+portString, router)
}
