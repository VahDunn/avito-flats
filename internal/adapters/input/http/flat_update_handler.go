package http

import (
	"avito-flats/internal/domain/valueobjects"
	"avito-flats/internal/usecases"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strconv"
)

// структура ручки, того, что она принимает и типового ответа на ошибку - POST запрос на обновление статуса квартиры )
type (
	UpdateFlatStatusHandler struct {
		UpdateFlatStatusUsecase usecases.UpdateFlatStatusUsecase
	}

	UpdateFlatStatusIn struct {
		FlatID           string
		ModerationStatus string
	}
)

// конструктор
func NewUpdateFlatStatusHandler(usecase usecases.UpdateFlatStatusUsecase) UpdateFlatStatusHandler {
	return UpdateFlatStatusHandler{UpdateFlatStatusUsecase: usecase}
}

func (u *UpdateFlatStatusHandler) updateFlatStatus(w http.ResponseWriter, r *http.Request) {
	requestID := uuid.New().String()

	userType := r.Context().Value("userType")

	// Только для модераторов
	if userType != valueobjects.Moderator {
		sendErrorResponse(w, "только для модераторов", requestID, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", requestID, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var in UpdateFlatStatusIn
	if err := json.Unmarshal(body, &in); err != nil {
		sendErrorResponse(w, "Error unmarshalling JSON", requestID, http.StatusBadRequest)
		return
	}

	flatID, err := strconv.ParseInt(in.FlatID, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing flatID", requestID, http.StatusBadRequest)
		return
	}

	tmp_status, err := strconv.ParseInt(in.ModerationStatus, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing ModerationStatus", requestID, http.StatusBadRequest)
		return
	}

	status := valueobjects.ModerationStatus(tmp_status)

	if status != 0 && status != 1 && status != 2 && status != 3 {
		sendErrorResponse(w, "Invalid ModerationStatus value", requestID, http.StatusBadRequest)
		return
	}

	updatedFlat, err := u.UpdateFlatStatusUsecase.UpdateFlatStatus(flatID, status)
	if err != nil {
		sendErrorResponse(w, "Error updating flat status", requestID, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(updatedFlat)
	if err != nil {
		sendErrorResponse(w, "Error marshaling flat", requestID, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
