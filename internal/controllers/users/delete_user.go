package users

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"

	"api_user/internal/services/users"

	"github.com/go-chi/chi/v5"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		logrus.Errorf("Error retrieving the User ID: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = users.DeleteUser(userID)
	if err != nil {
		logrus.Errorf("Error deleting the User: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
