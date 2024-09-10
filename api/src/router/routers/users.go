package routers

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnFollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/followers",
		Method:   http.MethodGet,
		Function: controllers.GetFollowersByUser,
		NeedAuth: true,
	},

	{
		URI:      "/users/{userId}/following",
		Method:   http.MethodGet,
		Function: controllers.GetFollowingByUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{userId}/update-password",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		NeedAuth: true,
	},
}
