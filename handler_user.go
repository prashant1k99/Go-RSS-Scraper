package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/prashant1k99/Go-RSS-Scraper/internal/database"
)

func (apiCfg apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := apiCfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid user id")
		return
	}

	user, err := apiCfg.DB.GetUserById(r.Context(), userUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
