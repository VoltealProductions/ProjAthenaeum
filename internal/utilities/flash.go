package utilities

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
)

func SetFlash(w http.ResponseWriter, name string, value []byte, path string) {
	c := &http.Cookie{Name: name, Value: encode(value), Path: path}
	http.SetCookie(w, c)
}

func GetFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}

	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}

	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return value, nil
}

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

func GetFlashMessage(w http.ResponseWriter, r *http.Request) (string, string) {
	sfm := getSuccessFm(w, r)
	if sfm != "" {
		return "success", sfm
	}

	efm := getErrorFm(w, r)
	if efm != "" {
		return "error", efm
	}

	return "", ""
}

func getSuccessFm(w http.ResponseWriter, r *http.Request) string {
	fm, err := GetFlash(w, r, "success")
	if err != nil {
		logger.LogErr(err.Error(), 500)
		return ""
	}

	s := string(fm)

	if s == "" {
		return ""
	}

	return s
}

func getErrorFm(w http.ResponseWriter, r *http.Request) string {
	fm, err := GetFlash(w, r, "error")
	if err != nil {
		logger.LogErr(err.Error(), 500)
		return ""
	}

	s := string(fm)

	if s == "" {
		return ""
	}

	return s
}
