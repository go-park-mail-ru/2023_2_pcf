FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o apiserver ./cmd/apiserver/main.go

FROM debian AS apiserver

WORKDIR /app
COPY --from=builder /app/apiserver ./
COPY ./configs ./configs
COPY ./File ./File

CMD ["./apiserver"]
