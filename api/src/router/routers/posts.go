package routers

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.CreatePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts",
		Method:   http.MethodGet,
		Function: controllers.GetPosts,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}",
		Method:   http.MethodGet,
		Function: controllers.GetPostById,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}",
		Method:   http.MethodPut,
		Function: controllers.UpdatePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}",
		Method:   http.MethodDelete,
		Function: controllers.DeletePost,
		NeedAuth: true,
	}, {
		URI:      "/users/{userId}/posts",
		Method:   http.MethodGet,
		Function: controllers.GetPostsByUser,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}/like",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{postId}/unlike",
		Method:   http.MethodPost,
		Function: controllers.UnlikePost,
		NeedAuth: true,
	},
}
