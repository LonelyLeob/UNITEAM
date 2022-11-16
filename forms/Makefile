.PHONY:build
build:
	go build -o forms -v ./cmd/main/...

.DEFAULT_GOAL:build

.PHONY:up
up:
	migrate -path migrations -database postgres://user@localhost:5432/forms up

.PHONY:down
down:
	migrate -path migrations -database postgres://user@localhost:5432/forms down