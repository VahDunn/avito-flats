package http

import (
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// UpdateFlatStatusHandler отвечает за обработку запросов на обновление статуса квартиры модератором.
type UpdateFlatStatusHandler struct {
	UpdateFlatStatusUsecase usecases.UpdateFlatStatusUsecase
}

type UpdateFlatStatusIn struct {
	FlatID           string
	ModerationStatus string
}

// UpdateFlatStatusHandler конструктор для создания обработчика
func NewUpdateFlatStatusHandler(usecase usecases.UpdateFlatStatusUsecase) UpdateFlatStatusHandler {
	return UpdateFlatStatusHandler{UpdateFlatStatusUsecase: usecase}
}

// updateFlatStatus обрабатывает POST-запросы по пути /flat/update
func (u *UpdateFlatStatusHandler) updateFlatStatus(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var in UpdateFlatStatusIn
	if err := json.Unmarshal(body, &in); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	flatID, err := strconv.ParseInt(in.FlatID, 10, 64)
	if err != nil {
		http.Error(w, "Error parsing flatID", http.StatusBadRequest)
		return
	}

	tmp_status, err := strconv.ParseInt(in.ModerationStatus, 10, 64)
	if err != nil {
		http.Error(w, "Error parsing ModerationStatus", http.StatusInternalServerError)
		return
	}

	status := valueobjects.ModerationStatus(tmp_status)

	if status != 0 && status != 1 && status != 2 && status != 3 {
		http.Error(w, "Invalid ModerationStatus value", http.StatusBadRequest)
		return
	}

	updatedFlat, err := u.UpdateFlatStatusUsecase.UpdateFlatStatus(flatID, status)
	if err != nil {
		http.Error(w, "Error creating house", http.StatusInternalServerError)
		return
	}

	// Сериализуем параметры дома в JSON
	response, err := json.Marshal(updatedFlat)
	if err != nil {
		http.Error(w, "Error marshaling flat", http.StatusInternalServerError)
		return
	}

	// Настраиваем заголовки и отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
