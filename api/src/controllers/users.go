package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
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

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response.ERROR(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.Follow(userId, followerId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		response.ERROR(w, http.StatusForbidden, errors.New("you cannot unFollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if err = repo.UnFollow(userId, followerId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func GetFollowersByUser(w http.ResponseWriter, r *http.Request) {

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
	followers, err := repo.GetFollowersByUser(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, followers)

}

func GetFollowingByUser(w http.ResponseWriter, r *http.Request) {
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
	following, err := repo.GetFollowingByUser(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, following)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := auth.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	userIdOnParams := mux.Vars(r)
	userId, err := strconv.ParseUint(userIdOnParams["userId"], 10, 64)

	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
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

	var passwords models.Password
	if err = json.Unmarshal(bodyRequest, &passwords); err != nil {
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
	passwordOnBd, err := repo.GetPassword(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(passwords.Old, passwordOnBd); err != nil {
		response.ERROR(w, http.StatusUnauthorized, errors.New("the old password is incorrect"))
		return
	}

	passwordWithHash, err := security.Hash(passwords.New)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePassword(userId, string(passwordWithHash)); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
