# Стадия сборки для apiserver
FROM golang:latest AS apiserver-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/apiserver ./cmd/apiserver/main.go

# Финальный образ для apiserver
FROM debian AS apiserver
WORKDIR /app
COPY --from=apiserver-builder /app/apiserver ./apiserver
COPY ./configs ./configs
COPY ./File ./File
RUN chmod 77 ./apiserver
CMD ["./apiserver"]

