package middleware

import (
	"admin/db"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	conn *db.Conn
}

func NewAuthMiddleware(conn *db.Conn) *AuthMiddleware {
	return &AuthMiddleware{conn}
}

func (m *AuthMiddleware) CookieAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		session, err := m.conn.GetSessionByID(sessionCookie.Value)
		if err != nil {
			log.Println("[ERROR]:", err.Error())
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		if session == nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
