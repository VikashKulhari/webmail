package handlers

import (
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf=8")
	w.WriteHeader(statusCode)
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "error while  marshal", http.StatusInternalServerError)
		return
	}
	// cleanedResponse := strings.Replace(string(response), "\"", "'", -1)
	// httplog.LogEntrySetField(ctx,"Response Data", cleanedResponse)
	//after implementing context logging
	w.Write(response)
}
