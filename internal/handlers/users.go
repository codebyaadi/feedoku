package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/codebyaadi/rss-feed-agg/internal/auth"
	"github.com/codebyaadi/rss-feed-agg/internal/database"
	"github.com/codebyaadi/rss-feed-agg/internal/utils"
)

func (apiCfg *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error hashing password: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:           uuid.New(),
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "user created successfully",
		"success": true,
		"data":    convertDatabaseUserToAPIUser(user, "", ""),
	})
}

func (apiCfg *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting user: %v", err))
		return
	}

	if !utils.CheckPasswordHash(params.Password, user.PasswordHash) {
		utils.RespondWithError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	accessToken, refreshToken, err := auth.GenerateJWT(user.ID, user.Name, user.Email, user.ApiKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error generating tokens: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "user logged in successfully",
		"success": true,
		"data":    convertDatabaseUserToAPIUser(user, accessToken, refreshToken),
	})
}

func (apiCfg *Handler) GetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "user retrieved successfully",
		"success": true,
		"data":    convertDatabaseUserToAPIUser(user, "", ""),
	})
}

func (apiCfg *Handler) GetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "posts retrieved successfully",
		"success": true,
		"data":    convertDatabasePostsToAPIPosts(posts),
	})
}
