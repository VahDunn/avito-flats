package http

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// HouseHandler отвечает за обработку запросов на создание дома.
type CreateFlatHandler struct {
	CreateFlatUsecase usecases.CreateFlatUsecase
}
type CreateFlatIn struct {
	houseid   string `json:"house_id"`
	number    string `json:"number"`
	price     string `json:"price"`
	roomcount string `json:"room_count"`
}

func NewFlatHandler(usecase usecases.CreateFlatUsecase) CreateFlatHandler {
	return CreateFlatHandler{CreateFlatUsecase: usecase}
}

// createNewFlat обрабатывает POST-запросы по пути /flat/create
func (h *CreateFlatHandler) createNewFlat(w http.ResponseWriter, r *http.Request) {
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

	var in CreateFlatIn
	if err := json.Unmarshal(body, &in); err != nil {
		sendErrorResponse(w, "Error unmarshalling JSON", requestID, http.StatusBadRequest)
		return
	}

	houseID, err := strconv.ParseInt(in.houseid, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Invalid house id", requestID, http.StatusBadRequest)
		return
	}

	number, err := strconv.ParseInt(in.number, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Invalid flat number", requestID, http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseInt(in.price, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Invalid price", requestID, http.StatusBadRequest)
		return
	}

	roomCount, err := strconv.ParseInt(in.roomcount, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Invalid room count", requestID, http.StatusBadRequest)
		return
	}

	newFlat, err := h.CreateFlatUsecase.CreateNewFlat(entities.HouseID(houseID), number, price, roomCount)
	if err != nil {
		// Здесь можно добавить логирование ошибки
		sendErrorResponse(w, "Error creating flat", requestID, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(newFlat)
	if err != nil {
		sendErrorResponse(w, "Error marshaling flat", requestID, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // Используем 201 Created для успешного создания ресурса
	w.Write(response)
}
