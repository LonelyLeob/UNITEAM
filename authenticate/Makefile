.PHONY:build
build:
	go build -o server -v ./cmd/main/...

.DEFAULT_GOAL:build

.PHONY:up
up:
	migrate -path migrations -database postgres://user@localhost:5432/users up

.PHONY:down
down:
	migrate -path migrations -database postgres://user@localhost:5432/users down