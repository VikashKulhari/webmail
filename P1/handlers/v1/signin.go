package v1

import (
	"encoding/json"
	"net/http"

	"github.com/VikashKulhari/P1/entities"
)

func (h *handlerV1) SignIn(w http.ResponseWriter, r *http.Request) {
	var req entities.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	jwtString, err1 := h.Service.SignIn(req)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusConflict)
		// if err == "userNotFound" {
		// 	http.Error(w, "User Not Signed Up...Sign In First", http.StatusUnauthorized)
		// }
		// if err.Error() == "incorrectPassword" {
		// 	http.Error(w, "Passowrd Does Not Match", http.StatusUnauthorized)
		// } else {
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// }
	}
	json.NewEncoder(w).Encode(map[string]string{"token": jwtString})
	// w.WriteHeader(http.StatusOK)

}