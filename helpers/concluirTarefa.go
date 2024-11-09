package helpers

import (
	"encoding/json"
	"net/http"
)

func ConcluirTarefa(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.URL.Query().Get("id")

		for i, tarefa := range Tarefas {
			if tarefa.ID == idStr {
				Tarefas[i].Concluida = true
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(Tarefas[i])
				return
			}
		}

		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
