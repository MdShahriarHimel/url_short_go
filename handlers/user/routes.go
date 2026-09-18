package user

import (
	"net/http"
	"url_short/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager middleware.Manager) {
	mux.Handle(
		"POST /users/login",
		manager.AppendMiddleWare(
			nil,
			http.HandlerFunc(h.LoginUser),
		),
	)

	mux.Handle(
		"POST /users/register",
		manager.AppendMiddleWare(
			nil,
			http.HandlerFunc(h.RegisterUser),
		),
	)

	mux.Handle(
		"GET /users",
		manager.AppendMiddleWare(
			nil,
			http.HandlerFunc(h.GetUserList),
		),
	)

	mux.Handle(
		"POST /users/refresh",
		manager.AppendMiddleWare(
			nil,
			http.HandlerFunc(h.GetNewAccessToken),
		),
	)

	mux.Handle(
		"POST /users/logout",
		manager.AppendMiddleWare(
			[]middleware.MiddleWare{middleware.AuthJwt},
			http.HandlerFunc(h.LogOut),
		),
	)

}
