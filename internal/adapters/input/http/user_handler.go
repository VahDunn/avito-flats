package http

import (
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type UserHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) UserHandler {
	return UserHandler{userUsecase: usecase}
}

func (h *UserHandler) DummyLogin(w http.ResponseWriter, r *http.Request) {

	requestID := uuid.New().String()
	userType_tmp := r.URL.Query().Get("user_type")

	userID := requestID
	fmt.Println(userType_tmp)

	if userType_tmp != "client" && userType_tmp != "moderator" {
		sendErrorResponse(w, "Invalid user type. Must be 'client' or 'moderator'", requestID, http.StatusBadRequest)
		return
	}

	var userType valueobjects.UserType
	if userType_tmp == "moderator" {
		userType = valueobjects.Moderator
	}
	if userType_tmp == "client" {
		userType = valueobjects.Client
	}

	newUser, err := h.userUsecase.DummyLogin(userType, userID)
	if err != nil {
		sendErrorResponse(w, "Error creating user", requestID, http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(newUser)
	if err != nil {
		sendErrorResponse(w, "Error marshaling user", requestID, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
