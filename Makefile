NAME=url-shortener

default: run

build:
	@go build -o ${NAME}

dev:
	@go run ./main.go

run: build
	@./${NAME}
