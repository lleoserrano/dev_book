package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/response"
	"encoding/json"
	"errors"
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
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := auth.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

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

	repo := repository.NewPostsRepository(db)
	postSavedInDatabase, err := repo.GetPostById(postId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if postSavedInDatabase.AuthorID != userIdOnToken {
		response.ERROR(w, http.StatusForbidden, errors.New("You can only update your own posts"))
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

	if err = post.Prepare(); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePost(postId, post); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := auth.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

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

	repo := repository.NewPostsRepository(db)
	postSavedInDatabase, err := repo.GetPostById(postId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if postSavedInDatabase.AuthorID != userIdOnToken {
		response.ERROR(w, http.StatusForbidden, errors.New("You can only delete your own posts"))
		return
	}

	if err = repo.DeletePost(postId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewPostsRepository(db)
	posts, err := repo.GetPostsByUser(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, posts)
}
func LikePost(w http.ResponseWriter, r *http.Request) {
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
	if err = repo.LikePost(postId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
func UnlikePost(w http.ResponseWriter, r *http.Request) {
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
	if err = repo.UnlikePost(postId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
