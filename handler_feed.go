package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
		HandleSqlError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (apiCfg apiConfig) getFeedByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	query := r.URL.Query()
	limit := 10
	if li := query.Get("limit"); li != "" {
		newLimit, err := strconv.Atoi(li)
		if err != nil {
			fmt.Println("Unable to parse Limit")
		} else {
			limit = newLimit
		}
	}
	feeds, err := apiCfg.DB.GetFeedsByUser(r.Context(), database.GetFeedsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		HandleSqlError(w, err)
		return
	}

	convertedFeeds := make([]Feed, 0, len(feeds))
	for _, feed := range feeds {
		convertedFeed := databaseFeedToFeed(feed)
		convertedFeeds = append(convertedFeeds, convertedFeed)
	}

	respondWithJSON(w, http.StatusOK, convertedFeeds)
}
