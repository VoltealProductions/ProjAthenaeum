package handlers

import (
	"net/http"
	"time"

	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/session"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/hash"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages/system"
	"github.com/google/uuid"
)

func GetRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, system.RegisterIndex())
	if err != nil {
		logger.LogFatal(err.Error(), 500)
	}
}

func PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	values, errors := parseUserFormValidationAndValidate(r)
	if len(errors) > 0 {
		err := utilities.RenderView(w, r, system.RegisterError(values, errors))
		if err != nil {
			logger.LogErr(err.Error(), 500)
		}
	}
	err := models.CreateUser(values.Username, values.Email, values.Password, r.FormValue("public") == "on")
	if err != nil {
		dbError := map[string]string{}
		dbError["errors"] = "Unable to create user."
		err := utilities.RenderView(w, r, system.RegisterError(values, dbError))
		if err != nil {
			logger.LogErr(err.Error(), 500)
		}
	} else {
		fm := []byte("Your account was created successfully!")
		utilities.SetFlash(w, "success", fm, "/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func GetLoginPage(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, system.LoginIndex())
	if err != nil {
		logger.LogFatal(err.Error(), 500)
	}
}

func PostLoginPage(w http.ResponseWriter, r *http.Request) {
	values, errors := LoginFormValidate(r)
	if len(errors) > 0 {
		err := utilities.RenderView(w, r, system.LoginError(values, errors))
		if err != nil {
			logger.LogErr(err.Error(), 500)
		}
	}

	user, err := models.GetUserForLogin(values.Email)
	if err != nil {
		dbError := map[string]string{}
		dbError["errors"] = "Incorrect login credentials"
		err := utilities.RenderView(w, r, system.LoginError(values, dbError))
		if err != nil {
			logger.LogErr(err.Error(), 500)
		}
	}

	correctPass := hash.CheckPasswordHash(values.Password, user.Password)
	if !correctPass {
		logger.LogErr("the login credentials are incorrect", 500)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(time.Hour * 12)
	logger.LogInfo(sessionToken)

	session.Sessions[sessionToken] = session.Session{
		ID:       user.ID,
		Username: user.Username,
		Expiry:   expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		Expires:  expiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	})

	utilities.SetFlash(w, "success", []byte("You are now logged in!"), "/")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("session_token")
	sessionToken := c.Value

	delete(session.Sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})

	utilities.SetFlash(w, "success", []byte("You are now logged out!"), "/")
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	return vals, vals.RegValidate()
}

func LoginFormValidate(r *http.Request) (system.LoginFormValues, map[string]string) {
	r.ParseForm()

	vals := system.LoginFormValues{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	return vals, vals.LogValidate()
}
