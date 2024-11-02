package handlers

import (
	"fmt"
	"net/http"

	"github.com/codebyaadi/rss-feed-agg/internal/auth"
	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/codebyaadi/rss-feed-agg/internal/utils"
)

// AuthenticatedHandler defines a function type that includes the authenticated user.
type AuthenticatedHandler func(http.ResponseWriter, *http.Request, database.User)

// AuthMiddleware is a middleware that authenticates the user based on the API key.
func (h *Handler) AuthMiddleware(next AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := utils.GetAccessToken(r.Header)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid API key: %v", err))
			return
		}

		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "invalid token: "+err.Error())
			return
		}

		user, err := h.DB.GetUserByEmail(r.Context(), claims.Email)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "authentication failed")
			return
		}

		// Pass the authenticated user to the next handler
		next(w, r, user)
	}
}
