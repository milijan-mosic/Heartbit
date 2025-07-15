package articles

import (
	"encoding/json"
	"net/http"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HelloResponse{
		Message: "Single article!",
		Status:  http.StatusOK,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to encode JSON"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
