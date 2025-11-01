package handlers

import (
	"admin/components"
	"admin/db"
	"context"
	"log"
	"net/http"
)

type MainHandler struct {
	conn *db.Conn
}

func NewMainHandler(conn *db.Conn) *MainHandler {
	return &MainHandler{conn}
}

func (h *MainHandler) RenderDashboardHome(w http.ResponseWriter, r *http.Request) {
	if err := components.Home("/").Render(context.Background(), w); err != nil {
		log.Println("[ERROR]: Failed to render dashboard", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
