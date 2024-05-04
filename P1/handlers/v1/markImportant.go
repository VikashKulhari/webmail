package v1

import (
	"net/http"
	"strconv"

	ijwt "github.com/VikashKulhari/P1/jwt"
	"github.com/go-chi/chi/v5"
)

func (h *handlerV1) MarkImp(w http.ResponseWriter, r *http.Request) {
	mailID := chi.URLParam(r, "ID")
	intMailID, err := strconv.ParseUint(mailID, 10, 64)
	intMailID1 := uint(intMailID)
	if err != nil {
		w.Write([]byte(err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	email, ok := r.Context().Value(ijwt.ContextEmail).(string)

	if !ok {
		http.Error(w, "Email ID not found in context", http.StatusInternalServerError)
		return
	}
	err1 := h.Service.MarkImportant(intMailID1, email)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email Marked Important"))

}
