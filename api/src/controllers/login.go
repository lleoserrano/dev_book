package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	bodyResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyResponse, &user); err != nil {
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
	userOnBd, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(user.Password, userOnBd.Password); err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userOnBd.ID)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
}
