package helpers

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"log"
)

func CriarTarefa(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var novaTarefa TarefaModel
		err := json.NewDecoder(r.Body).Decode(&novaTarefa)
		if err != nil {
			http.Error(w, "Erro ao decodificar tarefa", http.StatusBadRequest)
			return
		}

		novaTarefa.ID = uuid.New().String()
		Tarefas = append(Tarefas, novaTarefa)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(novaTarefa)

		// Logando informações relevantes de forma clara
		log.Printf("Requisição recebida: %s %s\n", r.Method, r.URL.Path)
		log.Printf("Cabeçalhos: %v\n", r.Header)
		log.Printf("Corpo da requisição: %s\n", novaTarefa)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

