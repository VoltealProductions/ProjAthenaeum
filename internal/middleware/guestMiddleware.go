package middleware

import (
	"fmt"
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
)

func GuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := models.IsNotLoggedIn(w, r)
		if res {
			utilities.SetFlash(w, "error", []byte(fmt.Sprintf("Error %v: You were redirected.\nThe page you tried to access is for guests only, you are already logged in, there is nothing for you here.", http.StatusUnauthorized)), "/")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}
