package main

import (
	"net/http"

	"github.com/prashant1k99/Go-RSS-Scraper/internal/auth"
	"github.com/prashant1k99/Go-RSS-Scraper/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			HandleSqlError(w, err)
			return
		}

		handler(w, r, user)
	}
}
