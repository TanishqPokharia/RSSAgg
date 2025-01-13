package handlers

import (
	"RSSAggregator/internal/database"
	"RSSAggregator/middlewares"
	"RSSAggregator/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"net/http"
)

func HandlerCreateFeedFollow() middlewares.AuthenticatedHandler {
	return func(w http.ResponseWriter, r *http.Request, user *database.User, config *utils.ApiConfig) {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		createFeedFollowParams := database.CreateFeedFollowParams{
			UserID: user.ID,
		}
		err := decoder.Decode(&createFeedFollowParams)
		if err != nil {
			utils.ErrResponse(w, http.StatusBadRequest, fmt.Sprintf("Could not parse json: %s", err.Error()))
			return
		}
		feedFollow, err := config.DB.CreateFeedFollow(r.Context(), createFeedFollowParams)
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
			return
		}
		utils.JSONResponse(w, http.StatusOK, feedFollow)
	}
}

func HandlerGetFeedFollows() middlewares.AuthenticatedHandler {
	return func(w http.ResponseWriter, r *http.Request, user *database.User, config *utils.ApiConfig) {
		feedFollows, err := config.DB.GetFeedFollows(r.Context(), user.ID)
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
			return

		}
		utils.JSONResponse(w, http.StatusOK, feedFollows)
	}
}

func HandlerDeleteFeedFollow() middlewares.AuthenticatedHandler {
	return func(w http.ResponseWriter, r *http.Request, user *database.User, config *utils.ApiConfig) {
		feedFollowId, err := uuid.Parse(chi.URLParam(r, "feedFollowId"))
		if err != nil {
			utils.ErrResponse(w, http.StatusBadRequest, "Invalid feed id")
			return
		}
		err = config.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
			UserID: user.ID,
			ID:     feedFollowId,
		})
		if err != nil {
			var pgErr *pq.Error
			if errors.As(err, &pgErr) {
				utils.PgErrResponse(w, http.StatusBadRequest, utils.PgErrDesc{Message: pgErr.Message, Code: pgErr.Code.Name()})
				return
			}
			utils.ErrResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.JSONResponse(w, http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "deletion successful",
		})

	}
}
