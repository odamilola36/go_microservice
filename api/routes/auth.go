package routes

import (
	"microservices/api/resthandlers"
	"net/http"
)

func NewAuthRoutes(handler resthandlers.AuthHandlers) []*Route {
	return []*Route{
		{
			Method:       http.MethodPost,
			Path:         "/signup",
			HandlerFunc:  handler.SignUp,
			AuthRequired: false,
		},
		{
			Method:       http.MethodPut,
			Path:         "/user/{id}",
			HandlerFunc:  handler.UpdateUser,
			AuthRequired: true,
		},
		{
			Method:       http.MethodGet,
			Path:         "/get-user/{id}",
			HandlerFunc:  handler.GetUser,
			AuthRequired: true,
		},
		{
			Method:       http.MethodGet,
			Path:         "/get-users",
			HandlerFunc:  handler.GetAllUsers,
			AuthRequired: true,
		},
		{
			Method:       http.MethodDelete,
			Path:         "/delete-user/{id}",
			HandlerFunc:  handler.DeleteUser,
			AuthRequired: true,
		},
		{
			Method:       http.MethodPost,
			Path:         "/auth",
			HandlerFunc:  handler.SignIn,
			AuthRequired: false,
		},
	}
}
