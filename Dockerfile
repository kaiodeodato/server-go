# Etapa 1: Build (compilação do código Go)
FROM golang:1.23.3-alpine AS builder

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos de dependências go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixa as dependências do Go
RUN go mod tidy

# Copia o código fonte do seu projeto para o contêiner
COPY . .

# Compila o código Go e gera um binário chamado 'servidor'
RUN go build -o servidor .

# Etapa 2: Imagem final - a partir da imagem Alpine, que é mais leve
FROM alpine:latest

# Copia o binário compilado da etapa 1 para o contêiner final
COPY --from=builder /app/servidor /usr/local/bin/servidor

# Copia os templates e arquivos estáticos para o local onde o servidor Go os encontrará
# IMPORTANTE: O caminho precisa ser correto!
COPY ./templates /app/templates
COPY ./static /app/static
COPY .env /app/.env
# Expõe a porta 8080 para que o contêiner possa receber requisições
EXPOSE 8080

# Define o comando a ser executado quando o contêiner for iniciado
CMD ["/usr/local/bin/servidor"]
