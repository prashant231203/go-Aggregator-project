package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/prashant231203/go-Aggregator-project/internal/database"
)

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		name string "name"
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:", err))
		return
	}

	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
	})

	respondWithJSON(w, 200, struct{}{})
}
