package users

import (
	"api_user/internal/helpers"
	"api_user/internal/models"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Name, &data.Email, &data.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func CreateUser(user models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Error opening the database: %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)",
		user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		logrus.Errorf("Error inserting the user into the database: %s", err.Error())
		return err
	}

	return nil
}

func DeleteUser(userID uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("Error opening the database: %s", err.Error())
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		logrus.Errorf("Error deleting the user from the database: %s", err.Error())
		return err
	}

	return nil
}
