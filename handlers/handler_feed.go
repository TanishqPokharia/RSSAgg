package handlers

import (
	"RSSAggregator/internal/database"
	"RSSAggregator/middlewares"
	"RSSAggregator/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"net/http"
)

func HandlerCreateFeed() middlewares.AuthenticatedHandler {
	return func(w http.ResponseWriter, r *http.Request, u *database.User, config *utils.ApiConfig) {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		createFeedParams := database.CreateFeedParams{
			UserID: u.ID,
		}
		err := decoder.Decode(&createFeedParams)
		if err != nil {
			utils.ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Could not pass json: %s", err.Error()))
			return
		}

		feed, err := config.DB.CreateFeed(r.Context(), createFeedParams)
		if err != nil {
			var pgErr *pq.Error
			if errors.As(err, &pgErr) {
				utils.PgErrResponse(w, http.StatusBadRequest, utils.PgErrDesc{})
				return
			}
			utils.ErrResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.JSONResponse(w, http.StatusOK, feed)
	}
}

func HandlerGetFeeds() middlewares.AuthenticatedHandler {
	return func(w http.ResponseWriter, r *http.Request, user *database.User, config *utils.ApiConfig) {
		feeds, err := config.DB.GetFeeds(r.Context())
		if err != nil {
			var pgErr *pq.Error
			if errors.As(err, &pgErr) {
				utils.PgErrResponse(w, http.StatusBadRequest, utils.PgErrDesc{
					Message: pgErr.Message,
					Code:    pgErr.Code.Name(),
				})
				return
			}
			utils.ErrResponse(w, http.StatusInternalServerError, err.Error())
		}

		utils.JSONResponse(w, http.StatusOK, feeds)
	}
}
