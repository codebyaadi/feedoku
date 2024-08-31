package handlers

import (
	"fmt"
	"net/http"

	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/codebyaadi/rss-feed-agg/internal/utils"
)

// AuthenticatedHandler defines a function type that includes the authenticated user.
type AuthenticatedHandler func(http.ResponseWriter, *http.Request, database.User)

// AuthMiddleware is a middleware that authenticates the user based on the API key.
func (h *Handler) AuthMiddleware(next AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := utils.GetAPIKey(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid API key: %v", err))
			return
		}

		user, err := h.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "authentication failed")
			return
		}

		// Pass the authenticated user to the next handler
		next(w, r, user)
	}
}
