package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepared("creation"); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	user.ID, err = repo.CreateUser(user)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("nameOrNick"))

	if isTokenInvalid := auth.ValidateToken(r); isTokenInvalid != nil {
		response.ERROR(w, http.StatusUnauthorized, isTokenInvalid)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	users, err := repo.GetAll(nameOrNick)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	user, err := repo.GetById(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userIdOnToken, err := auth.ExtractUserId(r)

	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if userIdOnToken != userId {
		response.ERROR(w, http.StatusForbidden, errors.New("you can only update your own user"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepared("update"); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.UpdateUser(userId, user); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userIdOnToken, err := auth.ExtractUserId(r)

	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if userIdOnToken != userId {
		response.ERROR(w, http.StatusForbidden, errors.New("you can only delete your own user"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.DeleteUser(userId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
