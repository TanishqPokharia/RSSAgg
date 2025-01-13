package handlers

import (
	"RSSAggregator/internal/database"
	"RSSAggregator/utils"
	"errors"
	"github.com/lib/pq"
	"net/http"
)

func HandlerGetPosts(w http.ResponseWriter, r *http.Request, user *database.User, config *utils.ApiConfig) {
	posts, err := config.DB.GetPostsForUser(r.Context(), user.ID)
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
	utils.JSONResponse(w, http.StatusOK, posts)
}
