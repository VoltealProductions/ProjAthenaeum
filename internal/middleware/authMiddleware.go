package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/google/uuid"
)

func AuthedMiddleware(next http.Handler) http.Handler {
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
		userSession, exists := models.LoadSession(tkn)
		if !exists {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   "",
				Path:    "/",
				Expires: time.Now(),
			})
			utilities.SetFlash(w, "error", []byte(fmt.Sprintf("%v: You can not go here as a guest, please logged in!", http.StatusUnauthorized)), "/")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if userSession.IsExpired() {
			err := models.DeleteSession(tkn)
			if err != nil {
				utilities.SetFlash(w, "error", []byte(fmt.Sprintf("%v: Your session is expired, please logged in!", http.StatusUnauthorized)), "/")
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}

		newSessionToken := uuid.NewString()
		expiresAt := time.Now().Add(time.Hour * 12)
		models.StoreSesson(newSessionToken, userSession.ID, expiresAt)

		err = models.DeleteSession(tkn) // Delete the old session
		if err != nil {
			utilities.SetFlash(w, "error", []byte("Unable to delete session, you were already logged out!"), "/")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		// Set the new token as the users `session_token` cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    newSessionToken,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 12),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
		})

		next.ServeHTTP(w, r)
	})
}
