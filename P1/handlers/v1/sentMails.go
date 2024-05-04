package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	ijwt "github.com/VikashKulhari/P1/jwt"
)

func (h *handlerV1) GetSentMails(w http.ResponseWriter, r *http.Request) {
	email, ok := r.Context().Value(ijwt.ContextEmail).(string)
	fmt.Println(ok)
	if !ok {
		http.Error(w, "Email not found in context", http.StatusInternalServerError)
		return
	}
	fmt.Println(email)
	emails, err := h.Service.GetEmailsSentById(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email fetched Successfully"))
	json.NewEncoder(w).Encode(emails)

}
