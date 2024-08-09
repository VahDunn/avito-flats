package http

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// FlatsHandler отвечает за обработку запросов, связанных с квартирами.
type FlatsHandler struct {
	FlatsUsecase usecases.FlatsUsecase
}

func NewFlatsHandler(usecase usecases.FlatsUsecase) FlatsHandler {
	return FlatsHandler{FlatsUsecase: usecase}
}

// getFlats обрабатывает GET-запросы по пути /house/{id}.
func (h *FlatsHandler) getFlats(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из пути, далее идут нужные преобразования
	requestID := uuid.New().String()
	query := r.URL.Query()

	userTypeStr := query.Get("userType")
	userID := query.Get("userID")
	houseIDStr := query.Get("houseID")

	if userTypeStr == "" || userID == "" || houseIDStr == "" {
		sendErrorResponse(w, "Missing query parameters", requestID, http.StatusBadRequest)
		return
	}

	userTypeInt, err := strconv.Atoi(userTypeStr)
	if err != nil || (userTypeInt != int(valueobjects.Moderator) && userTypeInt != int(valueobjects.Client)) {
		sendErrorResponse(w, "Invalid userType value", requestID, http.StatusBadRequest)
		return
	}

	userType := valueobjects.UserType(userTypeInt)
	houseIDint, err := strconv.Atoi(houseIDStr)
	if err != nil {
		sendErrorResponse(w, "Invalid houseID value", requestID, http.StatusBadRequest)
		return
	}

	houseID := entities.HouseID(houseIDint)
	// Создаем структуру user_data для передачи данных в usecase
	user_data := entities.GetFlatsByHouseIDIn{
		UserType: userType,
		UserID:   userID,
		HouseID:  houseID,
	}

	if user_data.UserType != valueobjects.Moderator && user_data.UserType != valueobjects.Client {
		sendErrorResponse(w, "Unauthorized access", requestID, http.StatusUnauthorized)
		return
	}

	// Вызов usecase для получения списка квартир

	flats, err := h.FlatsUsecase.GetFlatsByHouseID(r.Context(), user_data)
	if err != nil {
		sendErrorResponse(w, "Error retrieving flats", requestID, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(flats)
	if err != nil {
		sendErrorResponse(w, "Error marshaling flat", requestID, http.StatusInternalServerError)
		return
	}

	// Настраиваем заголовки и отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
