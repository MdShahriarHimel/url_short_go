package middleware

import "net/http"

type MiddleWare func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewareList []MiddleWare
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewareList: make([]MiddleWare, 0),
	}
}

func (m *Manager) AppendMiddleWare(mid []MiddleWare, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next := handler

		for i := len(mid) - 1; i >= 0 && len(mid) != 0; i-- {
			next = mid[i](next)
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Manager) GlobalMiddlewareAssign(mid ...MiddleWare) {
	for _, middleware := range mid {
		m.globalMiddlewareList = append(m.globalMiddlewareList, middleware)
	}
}

func (m *Manager) GlobalMiddleWareWrapper(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next := mux

		for i := len(m.globalMiddlewareList) - 1; i >= 0 && len(m.globalMiddlewareList) != 0; i-- {
			next = m.globalMiddlewareList[i](next)
		}

		next.ServeHTTP(w, r)
	})
}
