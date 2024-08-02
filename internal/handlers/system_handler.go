package handlers

import (
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
	_, err := models.CreateUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		logger.LogErr(err.Error(), 503)
	}
}
