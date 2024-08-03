package handlers

import (
	"fmt"
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages/system"
)

func GetRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, system.Register())
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
		if err != nil {
			logger.LogErr(err.Error(), 503)
		}

		fm := []byte("Your account was created successfully!")
		utilities.SetFlash(w, "success", fm, "/")
		logger.LogInfo(fmt.Sprintf("Set A Cookie: %v. Redirecting!", fm))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		w.Write([]byte("NOT ACCEPTED"))
	}
}
