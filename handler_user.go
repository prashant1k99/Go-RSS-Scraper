package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
		HandleSqlError(w, err, "User")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (apiCfg apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limit, skip := getPagination(r)
	fmt.Println("Limit", limit, "skip", skip)
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
		Offset: skip,
	})
	fmt.Println("found posts:", len(posts))
	if err != nil {
		HandleSqlError(w, err, "Posts")
		return
	}
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
