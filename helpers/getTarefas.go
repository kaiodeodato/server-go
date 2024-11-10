package helpers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTarefas(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Defina um contexto com timeout para a operação
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := Client.Database("gotask_database").Collection("gotask_tickets")

    // Busque todas as tarefas
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Println("Erro ao buscar tarefas:", err)
        http.Error(w, "Erro ao buscar tarefas", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    // Armazene as tarefas em um slice
    var tarefas []TarefaModel
    if err := cursor.All(ctx, &tarefas); err != nil {
        log.Println("Erro ao decodificar tarefas:", err)
        http.Error(w, "Erro ao decodificar tarefas", http.StatusInternalServerError)
        return
    }

    log.Printf("Tarefas encontradas: %v", tarefas)

    if len(tarefas) == 0 {
        log.Println("Nenhuma tarefa encontrada")
    }
    
    json.NewEncoder(w).Encode(tarefas)
}
