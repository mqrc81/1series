package api

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/mqrc81/zeries/domain"
)

type UserHandler struct {
	store    domain.Store
	sessions *scs.SessionManager
}

func (h *UserHandler) Register() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}
