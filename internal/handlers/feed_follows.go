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

func (h *Handler) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	feedFollow, err := h.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating feed follow: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "follow feed created successfully",
		"success": true,
		"data":    models.ConvertDatabaseFeedFollowToAPIFeedFollow(feedFollow),
	})
}

func (h *Handler) GetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := h.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "feeds retrieved successfully",
		"success": true,
		"data":    models.ConvertDatabaseFeedFollowsToAPIFeedFollows(feeds),
	})
}

func (h *Handler) DeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := r.PathValue("feedFollowID")

	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing feed follow id: %v", err))
		return
	}

	err = h.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error parsing feed follow id: %v", err))
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "feed follow deleted successfully",
		"success": true,
	})
}
