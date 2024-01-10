package users

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"api_user/internal/models"
	"api_user/internal/services/users"

	"github.com/go-chi/chi/v5"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.FromString(chi.URLParam(r, "id"))

	if err != nil {
		logrus.Errorf("Error retrieving the user ID: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		logrus.Errorf("Error reading the request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = users.UpdateUser(userID, updatedUser)
	if err != nil {
		logrus.Errorf("Error during user update: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(updatedUser)
	_, _ = w.Write(response)
}
