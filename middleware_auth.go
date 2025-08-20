package main

import (
	"fmt"
	"net/http"

	"github.com/ABHINAVGUPTA02/rssagg/internal/auth"
	"github.com/ABHINAVGUPTA02/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := auth.GetAPIKEY(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get API key: ", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), key)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: ", err))
			return
		}

		handler(w, r, user)
	}
}
