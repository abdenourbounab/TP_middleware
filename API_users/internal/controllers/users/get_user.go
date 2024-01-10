package users

import (
	"api_user/internal/models"
	"api_user/internal/repositories/users"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetUser
// @Tags         Users
// @Summary      Get a user.
// @Description  Get a user.
// @Param        id           	path      string  true  "User UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value("userId").(uuid.UUID)

	user, err := users.GetUserById(userId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}
