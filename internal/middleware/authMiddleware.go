package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VoltealProductions/Athenaeum/internal/session"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/google/uuid"
)

func Test(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tkn := c.Value
		userSession, exists := session.Sessions[tkn]
		logger.LogInfo(fmt.Sprintf("Session exists: %v", exists))
		if !exists {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if userSession.IsExpired() {
			delete(session.Sessions, tkn)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		newSessionToken := uuid.NewString()
		expiresAt := time.Now().Add(time.Hour * 12)
		session.Sessions[newSessionToken] = session.Session{
			ID:       userSession.ID,
			Username: userSession.Username,
			Expiry:   expiresAt,
		}

		// Delete the older session token
		delete(session.Sessions, tkn)

		// Set the new token as the users `session_token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    newSessionToken,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 12),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   true,
		})

		next.ServeHTTP(w, r)
	})
}
