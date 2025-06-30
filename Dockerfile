FROM golang:1.24
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o poliglotim_api ./cmd/main.go
CMD ["./poliglotim_api"]