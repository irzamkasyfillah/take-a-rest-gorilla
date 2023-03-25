package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/irzam/my-app/api/user/controller"
	"github.com/irzam/my-app/api/user/middleware/request"
)

type UserRoutes struct {
	controller controller.UserControllerInterface
}

type UserRoutesInterface interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetUserHistory(w http.ResponseWriter, r *http.Request)
}

func NewUserRoutes(controller controller.UserControllerInterface) UserRoutesInterface {
	return &UserRoutes{
		controller: controller,
	}
}

func (route *UserRoutes) GetUserHistory(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Get req body
	var body request.UserGetHistoryRequest
	json.NewDecoder(r.Body).Decode(&body)

	var fix interface{}
	data, err := route.controller.UserGetHistoryService(r.Context(), &body, &w)
	if data != nil {
		fix = data
	} else {
		fix = err
	}

	res, _ := json.Marshal(fix)
	w.Write(res)
}

func (route *UserRoutes) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Get req body
	var body request.UserGetAllRequest
	json.NewDecoder(r.Body).Decode(&body)

	var fix interface{}
	data, err := route.controller.UserGetAllService(r.Context(), &body, &w)
	if data != nil {
		fix = data
	} else {
		fix = err
	}

	res, _ := json.Marshal(fix)
	w.Write(res)
}

func (route *UserRoutes) GetUser(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Get the req params from the url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	req := request.UserGetOneRequest{ID: uint(id)}

	res, _ := json.Marshal(route.controller.UserGetByIdService(r.Context(), &req, &w))
	w.Write(res)
}

func (route *UserRoutes) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Get req body
	var body request.UserCreateRequest
	json.NewDecoder(r.Body).Decode(&body)

	res, _ := json.Marshal(route.controller.UserCreateService(r.Context(), &body, &w))
	w.Write(res)
}

func (route *UserRoutes) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Get request from params and body
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var body request.UserUpdateRequest
	json.NewDecoder(r.Body).Decode(&body)
	body.ID = uint(id)

	res, _ := json.Marshal(route.controller.UserUpdateService(r.Context(), &body, &w))
	w.Write(res)
}

func (route *UserRoutes) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Get the req params from the url
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	req := request.UserDeleteRequest{ID: uint(id)}

	res, _ := json.Marshal(route.controller.UserDeleteService(r.Context(), &req, &w))
	w.Write(res)
}
