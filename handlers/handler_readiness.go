package handlers

import (
	"RSSAggregator/utils"
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, struct{}{})
}
