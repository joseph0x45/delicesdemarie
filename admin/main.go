package main

import (
	"admin/db"
	"admin/handlers"
	"admin/middleware"
	"embed"
	"flag"
	"log"
	"net/http"
)

//go:embed static/*
var static embed.FS

//go:generate tailwindcss -i static/input.css -o static/styles.css -m

func main() {
	port := flag.String("port", "8080", "The port to start admin on")
	dbPath := flag.String("db", "admin.db", "The SQLite database file")
	newDB := flag.Bool("new-db", false, "Start with a fresh database")
	flag.Parse()
	conn, err := db.NewConn(*dbPath, *newDB)
	if err != nil {
		panic(err)
	}
	authHandler := handlers.NewAuthHandler(conn)
	authMiddleware := middleware.NewAuthMiddleware(conn)
	mainHandler := handlers.NewMainHandler(conn)
	mux := http.NewServeMux()

	mux.Handle("GET /", authMiddleware.CookieAuth(http.HandlerFunc(mainHandler.RenderDashboardHome)))
	mux.HandleFunc("GET /auth", authHandler.RenderAuthPage)

	mux.HandleFunc("POST /api/auth", authHandler.HandleLogin)

	mux.HandleFunc("GET /static/", http.FileServer(http.FS(static)).ServeHTTP)

	server := http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}
	log.Printf("App launched on  http://0.0.0.0:%s\n", *port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
