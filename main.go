package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Print("starting server...")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("defaulting to port %s", port)
	}

	dbUrl := os.Getenv("POSTGRES_URL")
	if dbUrl == "" {
		log.Fatal("POSTGRES_URL must be set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Can't connect to postgres database %v", err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatalf("Can't reach postgres database %v", err)
	}

	log.Println("Successfully connected to database")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}