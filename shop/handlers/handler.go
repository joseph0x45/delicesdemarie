package handlers

import (
	"context"
	"log"
	"net/http"
	"shop/components"
	"shop/db"
)

type Handler struct {
	conn *db.Conn
}

func NewHandler(conn *db.Conn) *Handler {
	return &Handler{
		conn: conn,
	}
}

func (h *Handler) RenderShopPage(w http.ResponseWriter, r *http.Request) {
	products, err := h.conn.GetAllProducts()
	if err != nil {
    log.Println("[ERROR]:", err.Error())
		http.Error(w, "Une Erreur s'est produite!", http.StatusInternalServerError)
    return
	}
	components.Index(products).Render(context.Background(), w)
}
