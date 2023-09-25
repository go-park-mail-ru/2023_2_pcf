.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: docker
docker:
	docker-compose down
	docker-compose up --build

.PHONY: docker-build
docker-build:
	docker-compose build

.DEFAULT_GOAL := docker