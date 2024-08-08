package http

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type SubscribeHandler struct {
	subscribeUsecase usecases.SubscribeUsecase
}

type SubscribeHandlerIn struct {
	houseid entities.HouseID `json:"house_id`
	email   string           `json:"email`
}

func NewSubscribeHandler(subscribeUsecase usecases.SubscribeUsecase) *SubscribeHandler {
	return &SubscribeHandler{subscribeUsecase: subscribeUsecase}
}

func (h *SubscribeHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	userType := r.Context().Value("userType")

	if userType != valueobjects.Moderator && userType != valueobjects.Client {
		sendErrorResponse(w, "Unauthorized access", requestID, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", requestID, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var in SubscribeHandlerIn

	if err := json.Unmarshal(body, &in); err != nil {
		sendErrorResponse(w, "Error unmarshalling JSON", requestID, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		sendErrorResponse(w, "Unauthorized access", requestID, http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SubscribeHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	return
	// Функция отписки не реализована. Подписка будет вечной.
}
