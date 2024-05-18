package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VikashKulhari/P1/entities"
)

func (h *handlerV1) SendMail(w http.ResponseWriter, r *http.Request) {
	var req entities.EmailReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	err = h.Service.SendEmail(req)
	fmt.Println(err)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Email Sent Successfully"))

}
