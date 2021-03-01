package middlewares

import (
	"github.com/gorilla/mux"
)

// Router is base router of app
var Router = mux.NewRouter().StrictSlash(true)

// AccountRouter is sub router handling accounts' stuff
var AccountRouter = Router.PathPrefix("/account").Subrouter()
