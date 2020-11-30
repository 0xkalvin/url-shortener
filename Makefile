.PHONY: build database dev down run url-shortener

NAME=url-shortener

default: run

all: database url-shortener

build:
	@go build -o ${NAME}

database:
	@docker-compose up -d dynamodb redis

dev:
	@go run ./main.go

down:
	@docker-compose down --rmi all -v --remove-orphans

run: build
	@./${NAME}

url-shortener:
	@docker-compose up url-shortener
