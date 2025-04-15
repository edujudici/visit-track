# Etapa 1: Build
FROM golang:1.24-alpine AS builder

# Instala as dependências
RUN apk add --no-cache git

# Define o diretório de trabalho
WORKDIR /app

# Copia o código fonte para dentro do container
COPY . .

# Baixa as dependências do Go
RUN go mod tidy

# Compila o código
RUN go build -o /bin/visit-track ./cmd/visit_track

# Etapa 2: Imagem final
FROM alpine:latest

# Instala o certificado SSL (para requisições HTTPS, se necessário)
RUN apk --no-cache add ca-certificates

# Copia o binário compilado da etapa anterior
COPY --from=builder /bin/visit-track /bin/visit-track

# Expõe a porta do servidor (8080 no caso)
EXPOSE 8080

# Comando para rodar o servidor
CMD ["/bin/visit-track"]
