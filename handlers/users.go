package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ortin779/blog-aggregator/helpers"
	"github.com/ortin779/blog-aggregator/store"
)

type UserHandler struct {
	DB *store.Queries
}

type createUserParameters struct {
	Name string `json:"name"`
}

func NewUserHandler(db *store.Queries) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func (usrHandl *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	usrParams := createUserParameters{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&usrParams)

	if err != nil {
		helpers.RespondWithError(w, 400, err.Error())
		return
	}

	user := store.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      usrParams.Name,
	}

	usr, err := usrHandl.DB.CreateUser(r.Context(), user)
	if err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	helpers.RespondWithJSON(w, 200, usr)
}

func (usrHandl *UserHandler) GetUserByApikey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")
	parts := strings.Split(apiKey, " ")
	if len(parts) < 2 {
		helpers.RespondWithError(w, 400, "invalid api key")
	}

	usr, err := usrHandl.DB.GetUserByApikey(r.Context(), parts[1])
	if err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	helpers.RespondWithJSON(w, 200, usr)
}
