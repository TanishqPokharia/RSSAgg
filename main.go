package main

import (
	"RSSAggregator/internal/database"
	"RSSAggregator/utils"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	dotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found")
	}

	db := os.Getenv("DB_URL")
	if db == "" {
		log.Fatal("DB_URL not found")
	}
	conn, err := sql.Open("postgres", db)
	if err != nil {
		log.Fatal("Can't connect to db")
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(conn)

	queries := database.New(conn)
	apiConf := utils.ApiConfig{DB: queries}

	fmt.Printf("Running on port %s\n", port)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowedOrigins:   []string{"http://*", "https://*"},
	}))
	v1Router := handlerV1Router(&apiConf)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
