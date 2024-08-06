package handlers

import (
	"net/http"

	"github.com/VoltealProductions/Athenaeum/internal/models"
	"github.com/VoltealProductions/Athenaeum/internal/utilities"
	"github.com/VoltealProductions/Athenaeum/internal/utilities/logger"
	"github.com/VoltealProductions/Athenaeum/internal/views/pages/characters"
)

func CharacterIndex(w http.ResponseWriter, r *http.Request) {
	err := utilities.RenderView(w, r, characters.CharactersIndex(models.IsLoggedIn(r)))
	if err != nil {
		logger.LogFatal(err.Error(), 500)
	}
}
