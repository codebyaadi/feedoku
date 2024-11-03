package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/codebyaadi/rss-feed-agg/internal/models"
	"github.com/codebyaadi/rss-feed-agg/internal/utils"
	"github.com/google/uuid"
)

func (h *Handler) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	feed, err := h.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "feed created successfully",
		"success": true,
		"data":    models.ConvertDatabaseFeedToAPIFeed(feed),
	})
}

func (h *Handler) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.DB.GetFeeds(r.Context())

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "feeds retrieved successfully",
		"success": true,
		"data":    models.ConvertDatabaseFeedsToAPIFeeds(feeds),
	})
}
