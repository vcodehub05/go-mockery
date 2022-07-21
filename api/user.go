package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go-mockery/user/model"
)

func Router(svc model.Service) *mux.Router {
	res := resource{service: svc}
	router := mux.NewRouter()
	router.HandleFunc("/create", res.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/list", res.GetAllUser).Methods("GET", "OPTIONS")
	return router

}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type resource struct {
	service model.Service
}

func (res resource) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Printf("Unable to decode the request body.  %v", err)
	}

	err = res.service.InsertUser(user)
	var report string
	if err != nil {
		report = "failed to create"
		json.NewEncoder(w).Encode(report)
	}

	message := "user created Successfully"

	json.NewEncoder(w).Encode(message)

}

func (res resource) GetAllUser(w http.ResponseWriter, r *http.Request) {

	users, err := res.service.GetUser()

	if err != nil {
		fmt.Printf("Unable to get all user %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)
}
