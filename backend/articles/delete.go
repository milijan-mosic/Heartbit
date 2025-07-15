package articles

import (
	"encoding/json"
	"net/http"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status := http.StatusNoContent

	response := HelloResponse{
		Message: "Article is deleted...",
		Status:  status,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to encode JSON"}`))
		return
	}

	w.WriteHeader(status)
	w.Write(jsonResponse)
}
