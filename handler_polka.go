package main

import (
	"encoding/json"
	"net/http"

	"github.com/eliird/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerWebhookPolka(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Data struct {
			UserID int `json:"user_id"`
		} `json:"data"`
		Event string `json:"event"`
	}
	params := parameters{}

	apiKey, err := auth.GetAPIToken(r.Header)
	if err != nil || apiKey != cfg.polkaAPI {
		respondWithError(w, 401, "Couldn't create user")
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "API KEY Mismatch")
	}
	if params.Event != "user.upgraded" {
		respondWithError(w, 204, "Couldn't create user")
		return
	}
	// set chirpyred status to be true
	user, err := cfg.DB.UpdateUserStatus(params.Data.UserID, true)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusNoContent, user)

}
