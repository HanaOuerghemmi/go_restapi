package users

import (
	"net/http"

	"github.com/HanaOuerghemmi/go_restapi/types"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {

}
