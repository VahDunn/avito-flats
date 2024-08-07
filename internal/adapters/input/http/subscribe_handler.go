package http

import (
	"avito-flats/internal/domain/entities"
	"avito-flats/internal/usecases"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// HouseHandler отвечает за обработку запросов на создание дома.
type SubscribeHandler struct {
	SubcribeUsecase usecases.SubscribeUsecase
}
type SubscribeHouse struct {
	houseid   entities.HouseID
	user_id   int64
	developer string
}

func NewSubscribeHandler(usecase usecases.HouseUsecase) HouseHandler {
	return HouseHandler{HouseUsecase: usecase}
}

// CreateHouse обрабатывает POST-запросы по пути /house/create
func (h *HouseHandler) subscribeOnHouse(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var in CreateHouseIn
	if err := json.Unmarshal(body, &in); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}
	buildyear, err := strconv.ParseInt(in.buildyear, 10, 64)
	if err != nil {
		http.Error(w, "Error parsing buildyear", http.StatusInternalServerError)
	}
	newHouse, err := h.HouseUsecase.CreateHouse(in.address, buildyear, in.developer)
	if err != nil {
		http.Error(w, "Error creating house", http.StatusInternalServerError)
		return
	}

	// Сериализуем параметры дома в JSON
	response, err := json.Marshal(newHouse)
	if err != nil {
		http.Error(w, "Error marshaling house", http.StatusInternalServerError)
		return
	}

	// Настраиваем заголовки и отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
