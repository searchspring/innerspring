package api

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"searchspring.com/innerspring/common"
	"searchspring.com/innerspring/google"
	"searchspring.com/innerspring/model"
	"searchspring.com/innerspring/projects"
)

var router *mux.Router

// Handler - check routing and call methods
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	if router == nil {
		r, err := CreateRouter()
		if err != nil {
			common.WriteError(w, 500, err.Error())
			return
		}
		router = r
	}
	router.ServeHTTP(w, r)
}

// CreateRouter publicldf so we can test it.
func CreateRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	apiSubRouter := router.PathPrefix("/api").Subrouter()
	projectHandler := projects.NewHandler()
	googleDAO := google.NewDAO(common.NewClient(&http.Client{}))
	apiSubRouter.HandleFunc("/projects", wrapSignedInUserCheck(googleDAO.CheckUserLoggedIn, projectHandler.GetProjects)).Methods(http.MethodGet)
	apiSubRouter.HandleFunc("/project/{id}", wrapSignedInUserCheck(googleDAO.CheckUserLoggedIn, projectHandler.GetProject)).Methods(http.MethodGet)
	return router, nil
}

func wrapSignedInUserCheck(checkUserLoggedIn func(authorizationToken string) (string, error), apiRequest func(w http.ResponseWriter, r *http.Request, u *model.User)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if email, err := checkUserLoggedIn(authorization); err != nil {
			common.WriteError(w, 403, err.Error())
			return
		} else {
			domain := strings.TrimSpace(os.Getenv("DOMAIN"))
			if domain != "" && !strings.HasSuffix(email, "@"+domain) {
				common.WriteError(w, 403, "must have "+domain+" email address to use this systsem")
				return
			}
			u := &model.User{
				Email:       email,
				AccessToken: authorization,
			}
			apiRequest(w, r, u)
		}
	}
}
