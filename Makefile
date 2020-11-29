.PHONY: build database dev run url-shortener

NAME=url-shortener

default: run

all: database url-shortener

build:
	@go build -o ${NAME}

database:
	@docker-compose up -d dynamodb redis

dev:
	@go run ./main.go

run: build
	@./${NAME}

url-shortener:
	@docker-compose up url-shortener
