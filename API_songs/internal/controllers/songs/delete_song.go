package songs

import (
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
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
func DeleteSong(w http.ResponseWriter, r *http.Request) {

	songID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération de l'ID du song : %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = songs.DeleteSong(songID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression du song : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
