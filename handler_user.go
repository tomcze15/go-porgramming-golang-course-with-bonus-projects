package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"tomcze15/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	errDecode := decoder.Decode(&params)

	if errDecode != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", errDecode))
		return
	}

	user, errNewUser := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if errNewUser != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create user: %s", errNewUser))
		return
	}

	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}
