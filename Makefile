.PHONY: build
build:
	go build -v ./cmd/loginsystem

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: prepare_db
prepare_db:
	docker pull postgres
	mkdir -p ${HOME}/docker/volumes/postgres
	docker run --rm --name pg-docker -e POSTGRES_PASSWORD=4321 -e POSTGRES_DB=loginsys_db -d -p 5432:5432 -v "${HOME}/docker/volumes/postgres:/var/lib/postgresql/data" postgres

.PHONY: run
run: prepare_db	build start_binary
	
.PHONY: start_binary
start_binary:
	./loginsystem

.DEFAULT_GOAL := build