package v1

import (
	"encoding/json"
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
	if err != nil {
		if len(req.To) == 0 {
			http.Error(w, "no reciepient", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Email Sent Successfully"))
	// sendResponse(w, http.StatusOK, map[string]interface{}{
	// 	"messageID": "ok",
	// 	"messgae":   "Email Sent Successfully",
	// })
}
