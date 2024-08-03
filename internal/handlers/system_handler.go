package handlers

import (
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages/system"
)

func GetRegisterPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := utilities.RenderView(w, r, system.RegisterIndex())
	if err != nil {
		logger.LogFatal(err.Error(), 500)
	}
}

func PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	values, errors := parseUserFormValidationAndValidate(r)
	if len(errors) > 0 {
		err := utilities.RenderView(w, r, system.Register(values, errors))
		if err != nil {
			logger.LogErr(err.Error(), 500)
		}
	}
	err := models.CreateUser(values.Username, values.Email, values.Password, r.FormValue("public") == "on")
	if err != "" {
	} else {
		fm := []byte("Your account was created successfully!")
		utilities.SetFlash(w, "success", fm, "/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func parseUserFormValidationAndValidate(r *http.Request) (system.CreateFormValues, map[string]string) {
	r.ParseForm()

	vals := system.CreateFormValues{
		Username:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		Public:    r.FormValue("public"),
		TosAccept: r.FormValue("accepttos"),
	}

	return vals, vals.Validate()
}
