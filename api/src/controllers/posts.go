package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := auth.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post.AuthorID = userIdOnToken

	if err = post.Prepare(); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewPostsRepository(db)
	post.ID, err = repo.CreatePost(post)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, post)
}
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := auth.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewPostsRepository(db)
	posts, err := repo.GetPosts(userIdOnToken)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}
func GetPostById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)
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

	repo := repository.NewPostsRepository(db)
	post, err := repo.GetPostById(postId)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, post)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {}
func DeletePost(w http.ResponseWriter, r *http.Request) {}
