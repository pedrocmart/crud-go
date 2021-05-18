package routes

import (
	"net/http"

	"github.com/pedrocmart/crud-go/api/controllers"
)

type UserRoutes interface {
	Routes() []*Route
}

type userRoutesImpl struct {
	userController controllers.UserController
}

func NewUserRoutes(userController controllers.UserController) *userRoutesImpl {
	return &userRoutesImpl{userController}
}

func (r *userRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/user",
			Method:  http.MethodPost,
			Handler: r.userController.PostUser,
		},
		{
			Path:    "/user",
			Method:  http.MethodGet,
			Handler: r.userController.GetUsers,
		},
		{
			Path:    "/user/{user_id}",
			Method:  http.MethodGet,
			Handler: r.userController.GetUser,
		},
		{
			Path:    "/user/{user_id}",
			Method:  http.MethodPut,
			Handler: r.userController.PutUser,
		},
		{
			Path:    "/user/{user_id}",
			Method:  http.MethodDelete,
			Handler: r.userController.DeleteUser,
		},
	}
}
