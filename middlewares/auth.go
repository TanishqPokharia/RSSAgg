package middlewares

import (
	"RSSAggregator/internal/auth"
	"RSSAggregator/internal/database"
	"RSSAggregator/utils"
	"fmt"
	"net/http"
)

// AuthHandler This is a type of function that requires a user struct that must be provided to it,
// in order to proceed with the further service
type AuthHandler func(w http.ResponseWriter, r *http.Request, user *database.User, config *utils.ApiConfig)

// AuthMiddleware Injects the user struct into the handler function
// Expects db connection and a handler function which requires user struct
func AuthMiddleware(config *utils.ApiConfig, handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract the api key from headers
		apiKey, err := auth.GetApiKey(&r.Header)
		if err != nil {
			utils.ErrResponse(w, http.StatusUnauthorized, fmt.Sprintf("Not Authorized"))
			return
		}

		// Get the user and provide it to the handler for further action
		user, err := config.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			utils.ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Could not get user with api key : %s", apiKey))
		}

		// inject the user into the handler and call it
		handler(w, r, &user, config)
	}
}
