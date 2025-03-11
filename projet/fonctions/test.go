package fonctions

import (
	"encoding/json"
	"net/http"
)

func CheckWordPressHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("site-url")
	err := IsWordPressSite(url)
	response := make(map[string]string)

	if err != nil {
		response["result"] = "Le site n'est pas un site WordPress"
	} else {
		response["result"] = "Le site est un site WordPress"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
