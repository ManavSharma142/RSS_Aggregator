package main

import (
	"fmt"
	"net/http"

	"github.com/ManavSharma142/RSS_Aggregator/internal/auth"
	"github.com/ManavSharma142/RSS_Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, er := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if er != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", er))
			return
		}

		handler(w, r, user)
	}
}