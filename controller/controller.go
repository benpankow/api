package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api-commons/errors"
	"github.com/HackIllinois/api-rsvp/models"
	"github.com/HackIllinois/api-rsvp/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{id}/", alice.New().ThenFunc(GetUserRsvp)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(GetCurrentUserRsvp)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(CreateCurrentUserRsvp)).Methods("POST")
	router.Handle("/", alice.New().ThenFunc(UpdateCurrentUserRsvp)).Methods("PUT")
}

/*
	Endpoint to get the rsvp for a specified user
*/
func GetUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(rsvp)
}

/*
	Endpoint to get the rsvp for the current user
*/
func GetCurrentUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(rsvp)
}

/*
	Endpoint to create the rsvp for the current user
*/
func CreateCurrentUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var rsvp models.UserRsvp
	json.NewDecoder(r.Body).Decode(&rsvp)

	rsvp.ID = id

	err := service.CreateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_rsvp)
}

/*
	Endpoint to update the rsvp for the current user
*/
func UpdateCurrentUserRsvp(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("HackIllinois-Identity")

	if id == "" {
		panic(errors.UnprocessableError("Must provide id"))
	}

	var rsvp models.UserRsvp
	json.NewDecoder(r.Body).Decode(&rsvp)

	rsvp.ID = id

	err := service.UpdateUserRsvp(id, rsvp)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_rsvp, err := service.GetUserRsvp(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_rsvp)
}
