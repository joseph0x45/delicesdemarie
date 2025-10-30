package main

import (
	"context"
	"embed"
	"flag"
	"log"
	"net/http"
	"shop/components"
	"shop/db"
)

//go:embed static/*
var static embed.FS

//go:generate tailwindcss -i static/input.css -o static/styles.css -m

func main() {
	port := flag.String("port", "8080", "The port to start shop on")
	newDb := flag.Bool("new-db", false, "Start with a fresh database")
	dbPath := flag.String("db", "shop.db", "The SQLite database flie")
	flag.Parse()
	mux := http.NewServeMux()

	conn, err := db.NewConn(*dbPath, *newDb)
	if err != nil {
		panic(err)
	}

	if err := conn.Seed(); err != nil {
		panic(err)
	}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		components.Index().Render(ctx, w)
	})
	mux.HandleFunc("GET /static/", http.FileServer(http.FS(static)).ServeHTTP)

	server := http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}
	log.Printf("shop launched on port %s\n", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
