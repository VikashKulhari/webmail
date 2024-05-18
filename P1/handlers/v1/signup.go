package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VikashKulhari/P1/entities"
)

func (h *handlerV1) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		http.Error(w, "No Body", http.StatusBadRequest)
	}
	req := entities.User{}
	fmt.Println("r.body: ", r.Body)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid SignUp Request", http.StatusBadRequest)
	}
	fmt.Println("1: ", req.Email, "2: ", req.Password)
	err = h.Service.SignUp(req)
	if err != nil {
		http.Error(w, "Email already registered/ some error", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Signed Up Successfully"))
}
