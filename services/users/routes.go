package users

import (
	"fmt"
	"net/http"

	"github.com/HanaOuerghemmi/go_restapi/auth"
	"github.com/HanaOuerghemmi/go_restapi/types"
	"github.com/HanaOuerghemmi/go_restapi/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handlerLogin).Methods("POST")
	router.HandleFunc("/register", h.handlerRegister).Methods("POST")
}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload

	// parse payload JSON
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON payload: %v", err))
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request: %v", errors))
		return
	}

	// check if user exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("not found, invalid email or password lol1"))
		return
	}
	if user == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user doesnt exists"))
		return
	}

	// check password
	err = auth.ComparePassword(payload.Password, user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wrong password"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"login": "success"})
}
func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterPayload
	//? parse payload JSON

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSSON payload"))
		return
	}
	//? validate payload

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payloda validation error : %v", errors))
		return
	}
	//? check if user already exists
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}
	//? hashed password
	hashedPassword, err := auth.CreateHashedPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error hashing plain password"))
		return
	}

	//? register user
	err = h.store.CreateUser(types.User{
		UserName: payload.UserName,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating a new user : %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]bool{
		"register": true,
	})
}
