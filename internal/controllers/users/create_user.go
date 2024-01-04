package users

import (
	"api_user/internal/models"
	"api_user/internal/repositories/users"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		logrus.Errorf("Error reading the request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate a new UUID ID
	id, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("Error generating UUID identifier: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Assign the generated ID to the user
	newUser.Id = &id

	err = users.CreateUser(newUser)
	if err != nil {
		logrus.Errorf("Error creating the User: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(newUser)
	_, _ = w.Write(response)
}
