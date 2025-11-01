package handlers

import (
	"admin/db"
	"admin/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/oklog/ulid/v2"
)

type AuthHandler struct {
	conn *db.Conn
}

func NewAuthHandler(conn *db.Conn) *AuthHandler {
	return &AuthHandler{conn}
}

type payloadStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) RenderAuthPage(w http.ResponseWriter, r *http.Request) {}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	payload := &payloadStruct{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Println("[ERROR]: Failed to decode request body:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, err := h.conn.GetUserByUsername(payload.Username)
	if err != nil {
		log.Println("[ERROR]:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !utils.PasswordIsCorrect(payload.Password, user.Password) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionID := ulid.Make().String()
	err = h.conn.InsertSession(sessionID)
	if err != nil {
		log.Println("[ERROR]: Failed to create session:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(10 * 365 * 24 * time.Hour),
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
