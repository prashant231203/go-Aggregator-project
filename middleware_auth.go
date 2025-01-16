package main

import (
	"fmt"
	"net/http"

	"github.com/prashant231203/go-Aggregator-project/internal/auth"
	"github.com/prashant231203/go-Aggregator-project/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error getting user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
