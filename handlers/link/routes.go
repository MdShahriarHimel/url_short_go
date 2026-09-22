package link

import (
	"net/http"
	"url_short/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager middleware.Manager) {
	mux.Handle(
		"POST /links/create",
		manager.AppendMiddleWare(
			[]middleware.MiddleWare{middleware.AuthJwt},
			http.HandlerFunc(h.CreateShortLink),
		),
	)

	mux.Handle(
		"GET /links",
		manager.AppendMiddleWare(
			[]middleware.MiddleWare{middleware.AuthJwt},
			http.HandlerFunc(h.GetLinkList),
		),
	)

	mux.Handle(
		"GET /links/{link_id}",
		manager.AppendMiddleWare(
			[]middleware.MiddleWare{middleware.AuthJwt},
			http.HandlerFunc(h.GetLinkById),
		),
	)

	mux.Handle(
		"PUT /links/{link_id}",
		manager.AppendMiddleWare(
			[]middleware.MiddleWare{middleware.AuthJwt},
			http.HandlerFunc(h.UpdateLink),
		),
	)

	mux.Handle(
		"DELETE /links/{link_id}",
		manager.AppendMiddleWare(
			[]middleware.MiddleWare{middleware.AuthJwt},
			http.HandlerFunc(h.DeleteLink),
		),
	)

	mux.Handle(
		"GET /{short_code}",
		manager.AppendMiddleWare(
			nil,
			http.HandlerFunc(h.Redirect),
		),
	)
}
