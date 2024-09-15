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
	},
}
