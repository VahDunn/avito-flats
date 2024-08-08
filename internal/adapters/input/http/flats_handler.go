package http

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strconv"
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
	// Извлекаем ID из пути
	requestID := uuid.New().String()

	userType := r.Context().Value("userType")

	if userType != valueobjects.Moderator && userType != valueobjects.Client {
		sendErrorResponse(w, "Unauthorized access", requestID, http.StatusUnauthorized)
		return
	}
	// Преобразуем ID из строки в целое число
	id, err := strconv.Atoi(requestID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Вызываем usecase для получения списка квартир
	flats, err := h.FlatsUsecase.GetFlatsByHouseID(r.Context(), entities.HouseID(id))
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
