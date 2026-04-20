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

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.IUserRepository
}

func NewUserHandler(userDB database.IUserRepository) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

// Login godoc
// @Summary     User login
// @Description User login
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request  body  dto.GetJWTInput  true  "user credentials"
// @Success     200 {object} dto.GetJWTOutput
// @Failure     400 {object} Error
// @Failure     401 {object} Error
// @Failure     404 {object} Error
// @Failure     500 {object} Error
// @Router      /users/login   [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var u dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	user, err := h.UserDB.FindByEmail(u.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !user.ValidatePassword(u.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		error := Error{Message: "Unauthorized Error"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]any{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary     Create user
// @Description Create user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       request  body  dto.CreateUserInput  true  "user request"
// @Success     201
// @Failure     400 {object} Error
// @Failure     500 {object} Error
// @Router      /users   [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(&error)
		return
	}
	defer r.Body.Close()

	user, err := entity.NewUser(u.Name, u.Email, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(&error)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: "Internal Server Error"}
		json.NewEncoder(w).Encode(&error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
