package helpers

import (
	"github.com/google/uuid"
)

var Tarefas []TarefaModel

func init() {
	Tarefas = append(Tarefas, TarefaModel{ID: uuid.New().String(), Descricao: "Estudar Go", Concluida: false})
	Tarefas = append(Tarefas, TarefaModel{ID: uuid.New().String(), Descricao: "Ler um livro", Concluida: true})
	Tarefas = append(Tarefas, TarefaModel{ID: uuid.New().String(), Descricao: "escrever diário", Concluida: true})
}