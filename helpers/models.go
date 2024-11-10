package helpers

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TarefaModel struct {
	ID          string        `json:"id" bson:"_id"`                     // ID da tarefa, equivalente ao ObjectId do MongoDB
	Title       string        `json:"title" bson:"title"`                // Título da tarefa
	Description string        `json:"description" bson:"description"`    // Descrição da tarefa
	Status      string        `json:"status" bson:"status"`              // Status da tarefa (e.g., "in_progress")
	Priority    string        `json:"priority" bson:"priority"`          // Prioridade (e.g., "high")
	AssignedTo  string        `json:"assigned_to" bson:"assigned_to"`    // ID do usuário atribuído
	DueDate     time.Time     `json:"due_date" bson:"due_date"`          // Data de vencimento
	Comments    []Comentario  `json:"comments" bson:"comments"`          // Lista de comentários
	Subtasks    []Subtarefa   `json:"subtasks" bson:"subtasks"`          // Lista de subtarefas
	ActivityLog []LogAtividade `json:"activity_log" bson:"activity_log"` // Histórico de atividades
}

type LogAtividade struct {
	Action     string    `json:"action" bson:"action"`
	Timestamp  time.Time `json:"timestamp" bson:"timestamp"`
	UserID     string    `json:"user_id" bson:"user_id"`
	NewStatus  string    `json:"new_status,omitempty" bson:"new_status,omitempty"` // Opcional, usado para mudanças de status
}

type Subtarefa struct {
	Title  string `json:"title" bson:"title"`
	Status string `json:"status" bson:"status"`
}

type Comentario struct {
	UserID    string    `json:"user_id" bson:"user_id"`
	Content   string    `json:"content" bson:"content"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

// Estrutura para o Usuário
type User struct {
    ID      primitive.ObjectID `json:"id" bson:"_id"` // ID único do usuário no MongoDB
    Name    string             `json:"name" bson:"name"`
    Email   string             `json:"email" bson:"email"`
    Role    string             `json:"role" bson:"role"` // Ex: "admin", "member"
    TeamID  primitive.ObjectID `json:"team_id" bson:"team_id"` // ID da equipe (referência para a coleção de equipes)
}

type Team struct {
    ID        primitive.ObjectID   `json:"id" bson:"_id"`       // ID único da equipe no MongoDB
    Name      string               `json:"name" bson:"name"`     // Nome da equipe
    Members   []primitive.ObjectID `json:"members" bson:"members"` // Lista de IDs de membros (referências para a coleção `Users`)
    TaskIDs   []primitive.ObjectID `json:"task_ids" bson:"task_ids"` // Lista de IDs de tarefas associadas (referência para a coleção `Tasks`)
}