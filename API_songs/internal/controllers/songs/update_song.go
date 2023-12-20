package songs

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"songs/internal/models"
	"songs/internal/services/songs"
)

// GetCollection
// @Tags         songs
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Collection
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [get]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	songID, err := uuid.FromString(chi.URLParam(r, "id"))

	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du user : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedSong models.Songs
	err = json.NewDecoder(r.Body).Decode(&updatedSong)
	if err != nil {
		logrus.Errorf("Erreur lors de la lecture du corps de la requête : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = songs.UpdateSong(songID, updatedSong)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour du user : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(updatedSong)
	_, _ = w.Write(response)
}
