.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: docker
docker:
	docker-compose up --build

.PHONY: docker-build
docker-build:
	docker-compose build

.PHONY: test-cvg
test-cvg:
	go test ./... -coverprofile=сoverаge.out
	go tool cover -func coverage.out | grep total:

.PHONY: proto
proto:
	protoc -I proto --go_out=proto proto/auth.proto && protoc -I proto --go-grpc_out=proto proto/auth.proto && protoc -I proto --go_out=proto proto/select.proto && protoc -I proto --go-grpc_out=proto proto/select.proto

.DEFAULT_GOAL := docker