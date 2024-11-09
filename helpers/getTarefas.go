package helpers

import (
	"encoding/json"
	"net/http"
)

func GetTarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tarefas)
}
