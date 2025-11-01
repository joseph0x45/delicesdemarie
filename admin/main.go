package main
import (
  "context"
  "admin/components"
  "embed"
  "flag"
  "log"
  "net/http"
)

//go:embed static/*
var static embed.FS

//go:generate tailwindcss -i static/input.css -o static/styles.css -m

func main(){
  port := flag.String("port", "8080", "The port to start admin on")
  flag.Parse()
  mux := http.NewServeMux()

  mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    components.Index().Render(ctx, w)
  })
  mux.HandleFunc("GET /static/", http.FileServer(http.FS(static)).ServeHTTP)

  
  server := http.Server{
    Addr: ":"+ *port,
    Handler: mux,
  }
  log.Printf("admin launched on port %s\n", *port)
  if err := server.ListenAndServe(); err!= nil {
    panic(err)
  }
}

