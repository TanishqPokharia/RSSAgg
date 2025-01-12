package handlers

import (
	"RSSAggregator/internal/database"
	"RSSAggregator/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"net/http"
	"time"
)

func HandlerCreateUser(config *utils.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		u := database.User{}
		err := decoder.Decode(&u)
		if err != nil {
			utils.ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %v", err))
			return
		}
		user, err := config.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID:        uuid.New(),
			Name:      u.Name,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})

		if err != nil {
			var e *pq.Error
			if errors.As(err, &e) {
				utils.PgErrResponse(w, http.StatusBadRequest, utils.PgErrDesc{
					Message: e.Message,
					Code:    e.Code.Name(),
				})
				return
			}
			utils.ErrResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusCreated, user)
	}
}

// HandlerGetUserByApiKey simply sends the user as json response
func HandlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, u *database.User, config *utils.ApiConfig) {
	utils.JSONResponse(w, http.StatusOK, u)
}
