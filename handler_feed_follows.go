package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/prashant1k99/Go-RSS-Scraper/internal/database"
)

func (apiCfg apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	// Parse the request body
	type parameters struct {
		FeedID uuid.UUID `json:"feedId"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		HandleSqlError(w, err, "Feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feed))
}

func (apiCfg apiConfig) handlerGetAllFollowedFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	limit, skip := getPagination(r)

	followedFeeds, err := apiCfg.DB.GetAllFollowedFeeds(r.Context(), database.GetAllFollowedFeedsParams{
		UserID: user.ID,
		Limit:  limit,
		Offset: skip,
	})

	if err != nil {
		HandleSqlError(w, err, "Followed Feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowedToFeedFollowed(followedFeeds))
}

func (apiCfg apiConfig) handlerDeleteFollowedFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "followId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow id")
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		HandleSqlError(w, err, "Feed Follow")
		return
	}
	respondWithJSON(w, http.StatusAccepted, map[string]string{"message": "Unfollowed successfully"})
}
