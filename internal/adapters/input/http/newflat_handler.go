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
type NewFlatHandler struct {
	CreateFlatUsecase usecases.CreateFlatUsecase
}
type CreateFlatIn struct {
	houseid   string
	number    string
	price     string
	roomcount string
}

func CreateFlatHandler(usecase usecases.CreateFlatUsecase) NewFlatHandler {
	return NewFlatHandler{CreateFlatUsecase: usecase}
}

// createNewFlat обрабатывает POST-запросы по пути /flat/create
func (n *NewFlatHandler) createNewFlat(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var in CreateFlatIn
	if err := json.Unmarshal(body, &in); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	temp_houseid, err := strconv.ParseInt(in.houseid, 10, 64)
	if err != nil {
		http.Error(w, "Что-то здесь не так", http.StatusInternalServerError)
	}
	houseid := entities.HouseID(temp_houseid)

	number, err := strconv.ParseInt(in.number, 10, 64)
	if err != nil {
		http.Error(w, "Flat number error", http.StatusInternalServerError)
	}
	price, err := strconv.ParseInt(in.price, 10, 64)
	if err != nil {
		http.Error(w, "Price error", http.StatusInternalServerError)
	}
	roomcount, err := strconv.ParseInt(in.roomcount, 10, 64)
	if err != nil {
		http.Error(w, "Room count error", http.StatusInternalServerError)
	}

	newFlat, err := n.CreateFlatUsecase.CreateNewFlat(houseid, number, price, roomcount)
	if err != nil {
		http.Error(w, "Error creating flat", http.StatusInternalServerError)
		return
	}

	// Сериализуем параметры квартиры в JSON
	response, err := json.Marshal(newFlat)
	if err != nil {
		http.Error(w, "Error marshaling house", http.StatusInternalServerError)
		return
	}

	// Настраиваем заголовки и отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
