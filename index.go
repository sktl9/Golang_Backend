package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {

	var db *sql.DB
	var err error // err declared here for the first time in this scope

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	connStr := "user=" + DB_USER + " dbname=" + DB_NAME + " password=" + DB_PASSWORD + " host=" + DB_HOST + " port=" + DB_PORT + " sslmode=disable"

	db, err = sql.Open("postgres", connStr) // err used here, but not redeclared
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping() // err used here again
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Привет, мир!"))
	})

	// Определение порта для запуска сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}

	// Запуск сервера
	log.Printf("Сервер запущен на http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
