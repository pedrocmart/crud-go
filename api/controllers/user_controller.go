package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/pedrocmart/crud-go/api/models"
	"github.com/pedrocmart/crud-go/api/repository"
	"github.com/pedrocmart/crud-go/api/utils"
)

type UserController interface {
	PostUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	PutUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type userControllerImpl struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *userControllerImpl {
	return &userControllerImpl{userRepository}
}

func (c *userControllerImpl) PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{}
	err = json.Unmarshal(bytes, user)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = user.Validate()

	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	user, err = c.userRepository.Save(user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, user.ID))
	utils.WriteAsJson(w, user)
}

func (c *userControllerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {

	user, err := c.userRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, user)
}

func (c *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	user, err := c.userRepository.Find(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, user)
}

func (c *userControllerImpl) PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{}
	err = json.Unmarshal(bytes, user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user.ID = user_id

	err = user.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.userRepository.Update(user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.userRepository.Delete(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	// buildDeleteResponse(w, user_id)
	utils.WriteAsJson(w, "{}")
}
