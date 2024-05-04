package v1

import (
	"encoding/json"
	"net/http"

	"github.com/VikashKulhari/P1/entities"
)

func (h *handlerV1) SignUp(w http.ResponseWriter, r *http.Request) {
	var req entities.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid SignUp Request", http.StatusBadRequest)
	}

	err = h.Service.SignUp(req)
	if err != nil {
		http.Error(w, "Email already registered/ some error", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Signed Up Successfully"))
}
