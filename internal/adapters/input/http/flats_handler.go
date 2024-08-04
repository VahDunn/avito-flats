package http

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/usecases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
	idStr := chi.URLParam(r, "id")

	// Преобразуем ID из строки в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Вызываем usecase для получения списка квартир
	flats, err := h.FlatsUsecase.GetFlatsByHouseID(entities.HouseID(id))
	if err != nil {
		http.Error(w, "Error retrieving flats", http.StatusInternalServerError)
		return
	}

	// Сериализуем список квартир в JSON
	response, err := json.Marshal(flats)
	if err != nil {
		http.Error(w, "Error marshaling flats", http.StatusInternalServerError)
		return
	}

	// Настраиваем заголовки и отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
