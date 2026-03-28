package booking

import (
	"encoding/json"
	"net/http"
)



type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) HoldSeat(request *http.Request, response http.ResponseWriter) {
	// parse request
	var b Booking
	err:=json.NewDecoder(request.Body).Decode(&b)
	if err != nil {
		http.Error(response, "invalid request", http.StatusBadRequest)
		return
	}
	// call service
	err=h.Service.Book(b)
	if err != nil {
		http.Error(response, "failed to book seat", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("seat held successfully"))
}