package middleware

import (
	"errors"
	"net/http"

	"github.com/Taiwrash/go-rest/api"
	"github.com/Taiwrash/go-rest/internal/tools"

	// "github.com/btcsuite/btcd/database"
	log "github.com/sirupsen/logrus"
)

var UnAuthorisedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorisedError)
			api.RequestErrorHandler(w, UnAuthorisedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorisedError)
			api.RequestErrorHandler(w, UnAuthorisedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
