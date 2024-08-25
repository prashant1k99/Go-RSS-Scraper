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
		HandleSqlError(w, err, "Feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (apiCfg apiConfig) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit := 10
	skip := 0
	if li := query.Get("limit"); li != "" {
		newLimit, err := strconv.Atoi(li)
		if err != nil {
			fmt.Println("Unable to parse Limit")
		} else {
			limit = newLimit
		}
	}

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

// func (apiCfg apiConfig) getFeedByUser(w http.ResponseWriter, r *http.Request, user database.User) {
// 	query := r.URL.Query()
// 	limit := 10
// 	if li := query.Get("limit"); li != "" {
// 		newLimit, err := strconv.Atoi(li)
// 		if err != nil {
// 			fmt.Println("Unable to parse Limit")
// 		} else {
// 			limit = newLimit
// 		}
// 	}
// 	feeds, err := apiCfg.DB.GetFeedsByUser(r.Context(), database.GetFeedsByUserParams{
// 		UserID: user.ID,
// 		Limit:  int32(limit),
// 	})
// 	if err != nil {
// 		HandleSqlError(w, err, "Feed")
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
// }

// func (apiCfg apiConfig) getFeedById(w http.ResponseWriter, r *http.Request, user database.User) {
// 	feedId := chi.URLParam(r, "feedId")
// 	feedUUID, err := uuid.Parse(feedId)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid feedId")
// 		return
// 	}

// 	feed, err := apiCfg.DB.GetFeedById(r.Context(), feedUUID)
// 	if err != nil {
// 		HandleSqlError(w, err, "Feed")
// 		return
// 	}
// 	if feed.UserID != user.ID {
// 		respondWithError(w, http.StatusForbidden, "Unauthorized")
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
// }

// func (apiCfg apiConfig) updateFeedById(w http.ResponseWriter, r *http.Request, user database.User) {
// 	// Parse the request body
// 	type parameters struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	}
// 	decoder := json.NewDecoder(r.Body)

// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	// Validate the update feed params
// 	if params.Name == "" || params.Url == "" {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	feedId := chi.URLParam(r, "feedId")
// 	feedUUID, err := uuid.Parse(feedId)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid feedId")
// 		return
// 	}

// 	feed, err := apiCfg.DB.UpdateFeed(r.Context(), database.UpdateFeedParams{
// 		ID:        feedUUID,
// 		Name:      params.Name,
// 		UpdatedAt: time.Now().UTC(),
// 		Url:       params.Url,
// 		UserID:    user.ID,
// 	})
// 	if err != nil {
// 		HandleSqlError(w, err, "Feed")
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
// }

// func (apiCfg apiConfig) deleteFeedById(w http.ResponseWriter, r *http.Request, user database.User) {
// 	feedId := chi.URLParam(r, "feedId")
// 	feedUUID, err := uuid.Parse(feedId)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid feedId")
// 		return
// 	}

// 	feed, err := apiCfg.DB.DeleteFeed(r.Context(), database.DeleteFeedParams{
// 		ID:     feedUUID,
// 		UserID: user.ID,
// 	})
// 	if err != nil {
// 		HandleSqlError(w, err, "Feed")
// 		return
// 	}
// 	fmt.Println(feed)
// 	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
// }
