package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/prashant1k99/Go-RSS-Scraper/internal/database"
)

func (apiCfg apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	type parameters struct {
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	apiCfg.DB.CreateUsers(r.Context(), database.CreateUsersParams{
		ID:   uuid.New(),
		Name: params.Name,
	})

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
