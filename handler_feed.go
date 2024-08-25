package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prashant1k99/Go-RSS-Scraper/internal/database"
)

func (apiCfg apiConfig) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// Parse the request body
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		HandleSqlError(w, err, "Feed")
		return
	}

	_, err = apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		fmt.Println("Error creating follow feed after creating feed:", err)
	}

	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (apiCfg apiConfig) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	limit, skip := getPagination(r)

	feeds, err := apiCfg.DB.GetAllFeeds(r.Context(), database.GetAllFeedsParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	})
	if err != nil {
		HandleSqlError(w, err, "Feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}

func (apiCfg apiConfig) handlerBulkCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	// Parse the request body
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := []parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	for _, param := range params {
		_, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        uuid.New(),
			Name:      param.Name,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Url:       param.Url,
			UserID:    user.ID,
		})
		if err != nil {
			fmt.Println("Error creating feed:", err)
			return
		}
	}

	respondWithJSON(w, http.StatusOK, "Feeds created successfully")
}
