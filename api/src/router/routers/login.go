package routers

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute Route = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	NeedAuth: false,
}
