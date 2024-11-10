package main

import (
    "html/template"
    "log"
    "net/http"
    "codeGo/helpers"
    "path/filepath"
)

func main() {

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles(filepath.Join("/app", "templates", "index.html"))
        if err != nil {
            http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        tmpl.Execute(w, nil)
    })

    http.HandleFunc("/tarefas", helpers.GetTarefas)
    http.HandleFunc("/tarefa", helpers.CriarTarefa)

    log.Println("Servidor rodando na porta 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
