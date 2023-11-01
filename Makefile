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
	go tool cover -func coverage.out | grep total:
	rm coverage.out

.DEFAULT_GOAL := docker