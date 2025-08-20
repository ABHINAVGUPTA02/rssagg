package main

import (
	"fmt"
	"net/http"

	"github.com/ABHINAVGUPTA02/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	FeedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a feed follow: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowsToUserFeedFollows(FeedFollows))
}
