package main

import (
	"database/sql"
	"github.com/mike-kimani/fechronizo/v2/internal/feedfollows"
	"github.com/mike-kimani/fechronizo/v2/internal/feeds"
	"github.com/mike-kimani/fechronizo/v2/internal/middleware"
	"github.com/mike-kimani/fechronizo/v2/internal/scraper"
	"github.com/mike-kimani/fechronizo/v2/internal/users"
	"github.com/mike-kimani/fechronizo/v2/pkg/errors"
	"github.com/mike-kimani/fechronizo/v2/pkg/readiness"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mike-kimani/fechronizo/v2/internal/database"

	_ "github.com/lib/pq"
)

/* holds a connection to a database*/
type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in this environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in this environment")
	}

	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	router := chi.NewRouter()
	/* middleware configuration to allow connection to our server through a browser*/
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	db := database.New(conn)

	feedFollowsApiConfig := feedfollows.ApiConfig{
		DB: db,
	}

	usersApiConfig := users.ApiConfig{
		DB: db,
	}

	feedsApiConfig := feeds.ApiConfig{
		DB: db,
	}

	MiddlewareAuthApiConfig := middleware.ApiConfig{
		DB: db,
	}

	go scraper.StartScraping(db, 10, 1*time.Minute)

	v1Router := chi.NewRouter()
	/* Handler only fires on get requests*/
	v1Router.Get("/healthz", readiness.HandlerReadiness)
	v1Router.Get("/err", errors.HandlerErr)
	v1Router.Post("/users", usersApiConfig.HandlerCreateUser)
	v1Router.Get("/users", MiddlewareAuthApiConfig.MiddlewareAuth(usersApiConfig.HandlerGetUser))
	v1Router.Post("/feeds", MiddlewareAuthApiConfig.MiddlewareAuth(feedsApiConfig.HandlerCreateFeed))
	v1Router.Get("/feeds", feedsApiConfig.HandlerGetFeed)
	v1Router.Post("/followfeeds", MiddlewareAuthApiConfig.MiddlewareAuth(feedFollowsApiConfig.HandlerCreateFeedFollow))
	v1Router.Get("/followfeeds", MiddlewareAuthApiConfig.MiddlewareAuth(feedFollowsApiConfig.HandlerGetFeedFollows))
	v1Router.Delete("/followfeedstest", MiddlewareAuthApiConfig.MiddlewareAuth(feedFollowsApiConfig.HandlerDeleteFeedFollows))
	v1Router.Delete("/followfeeds/{feedFollowID}", MiddlewareAuthApiConfig.MiddlewareAuth(feedFollowsApiConfig.HandlerDeleteFeedFollows2))
	v1Router.Get("/userposts", MiddlewareAuthApiConfig.MiddlewareAuth(usersApiConfig.HandlerGetPostsForUser))

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
}
