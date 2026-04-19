package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/leoneville/goexpert/api/internal/dto"
	"github.com/leoneville/goexpert/api/internal/entity"
	"github.com/leoneville/goexpert/api/internal/infra/database"
)

type UserHandler struct {
	UserDB       database.IUserRepository
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(
	userDB database.IUserRepository, jwt *jwtauth.JWTAuth, jwtExpiresIn int,
) *UserHandler {
	return &UserHandler{UserDB: userDB, Jwt: jwt, JwtExpiresIn: jwtExpiresIn}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var u dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(u.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(u.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]any{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := entity.NewUser(u.Name, u.Email, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
