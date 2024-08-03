package handlers

import (
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages/system"
)

func GetRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, system.Register(utilities.GetFlashMessage(w, r)))
	if err != nil {
		logger.LogFatal(err.Error(), 500)
	}
}

func PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.FormValue("accepttos") == "true" {
		public := false
		if r.FormValue("public") == "true" {
			public = true
		}
		err := models.CreateUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"), public)
		if err != "" {
			utilities.SetFlash(w, "error", []byte(err), "/")
			http.Redirect(w, r, "/s/register", http.StatusSeeOther)
			return
		} else {
			fm := []byte("Your account was created successfully!")
			utilities.SetFlash(w, "success", fm, "/")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		utilities.SetFlash(w, "error", []byte("Accepting the terms and conditions is required to register!"), "/")
		http.Redirect(w, r, "/s/register", http.StatusSeeOther)
	}
}
