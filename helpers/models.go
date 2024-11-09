package helpers

type TarefaModel struct {
	ID        string    `json:"id"`
	Descricao string `json:"descricao"`
	Concluida bool   `json:"concluida"`
}