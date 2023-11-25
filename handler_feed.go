package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"github.com/dhruvsolanki0811/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConifg) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parrsing json %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      param.Name,
		Url:       param.URL,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConifg) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Coudn't Get feeds %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))

	
}
