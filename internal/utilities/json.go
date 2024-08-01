package utilities

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		logger.LogErr(fmt.Sprint("Resonding with 5xx error: ", msg), code)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errResponse{
		Error: msg,
	})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		logger.LogErr(fmt.Sprintf("Failed to marshal JSON response: %v\n", payload), 412)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(dat)
	if err != nil {
		log.Fatal(err)
	}
}
