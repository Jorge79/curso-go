package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jorge79/estudos-go/APIS/internal/dto"
	"github.com/Jorge79/estudos-go/APIS/internal/entity"
	"github.com/Jorge79/estudos-go/APIS/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB      database.UserInterface
	Jwt         *jwtauth.JWTAuth
	JwtEpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, JwtEpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:      userDB,
		Jwt:         jwt,
		JwtEpiresIn: JwtEpiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	r.Context().Value("token")
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtEpiresIn)).Unix(),
	})
	accesToken :=
		struct {
			AccesToken string `json:"acces_token"`
		}{
			AccesToken: token,
		}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accesToken)

}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
