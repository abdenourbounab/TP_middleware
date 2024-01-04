package users

import (
	"api_user/internal/models"
	repository "api_user/internal/repositories/users"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "User not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

func CreateUser(user models.User) error {
	err := repository.CreateUser(user)
	if err != nil {
		logrus.Errorf("Error during user creation: %s", err.Error())
		return err
	}
	return nil
}

func DeleteUser(userID uuid.UUID) error {
	err := repository.DeleteUser(userID)
	if err != nil {
		logrus.Errorf("Error during user deletion: %s", err.Error())
		return err
	}

	return nil
}

func UpdateUser(userID uuid.UUID, updatedUser models.User) error {
	user, err := repository.GetUserById(userID)
	if err != nil {
		logrus.Errorf("Error retrieving the user: %s", err.Error())
		return err
	}

	// Update the necessary fields of the retrieved user with the data from the updated user
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	err = repository.UpdateUser(user)
	if err != nil {
		logrus.Errorf("Error updating the user in the database: %s", err.Error())
		return err
	}

	return nil
}
