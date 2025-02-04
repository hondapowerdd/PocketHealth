package user

import (
	"encoding/json"
	"log"
	"net/http"
	"pockethealth/internchallenge/pkg/router"
	"strings"
)

// An UserApiController binds http requests to an api service and writes the service results to the http response
type UserApiController struct {
	service UserApiService
}

// NewUserApiController creates a new api controller
func NewUserApiController(s UserApiService) router.Router {
	return UserApiController{service: s}
}

// Routes returns all of the api route for the UserApiController
func (c UserApiController) Routes() router.Routes {
	return router.Routes{
		{
			Name:        "PostRegister",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/register",
			HandlerFunc: c.PostRegister,
		},
	}
}

type PostRegisterBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	// modify the request body to include favourite colour
	FavouriteColor string `json:"favouriteColour"`
}

type PostRegisterResponse struct {
	UserId string `json:"user_id"`
}

// PostRegister - Register a New User
func (c UserApiController) PostRegister(w http.ResponseWriter, r *http.Request) {
	// read request body
	data := &PostRegisterBody{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("request body: %+v\n", data)

	// call service
	userId, err := c.service.PostRegister(r.Context(), data.Name, data.Email, data.FavouriteColor) // modify the request body to include favourite colour
	if err != nil {
		panic(err)
	}

	// create and send response
	resp := PostRegisterResponse{
		UserId: userId,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		panic(err)
	}
}
