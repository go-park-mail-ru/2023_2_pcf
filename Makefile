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
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.DEFAULT_GOAL := docker