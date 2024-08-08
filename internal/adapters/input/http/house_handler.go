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
type HouseHandler struct {
	HouseUsecase usecases.HouseUsecase
}
type CreateHouseIn struct {
	Address   string `json:"address"`
	BuildYear string `json:"buildyear"`
	Developer string `json:"developer"`
}

func NewHouseHandler(usecase usecases.HouseUsecase) HouseHandler {
	return HouseHandler{HouseUsecase: usecase}
}

// CreateHouse обрабатывает POST-запросы по пути /house/create
func (h *HouseHandler) createNewHouse(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()
	userType := r.Context().Value("userType")

	if userType != valueobjects.Moderator {
		sendErrorResponse(w, "Только для модераторов", requestID, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", requestID, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var in CreateHouseIn
	if err := json.Unmarshal(body, &in); err != nil {
		sendErrorResponse(w, "Error unmarshalling JSON", requestID, http.StatusBadRequest)
		return
	}

	buildyear, err := strconv.ParseInt(in.BuildYear, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing buildyear", requestID, http.StatusBadRequest)
		return
	}

	newHouse, err := h.HouseUsecase.CreateHouse(in.Address, buildyear, in.Developer)
	if err != nil {
		sendErrorResponse(w, "Error creating house", requestID, http.StatusInternalServerError)
		return
	}

	SuccessNewHouseResponse(w, newHouse, requestID)
}

func sendJSONResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func SuccessNewHouseResponse(w http.ResponseWriter, data entities.House, requestID string) {
	response := data
	sendJSONResponse(w, response, http.StatusOK)
}
