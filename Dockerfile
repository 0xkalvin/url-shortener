# Build stage
FROM golang:1.14-alpine as builder

WORKDIR /url-shortener

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o binary

# Production stage
FROM alpine as production

WORKDIR /url-shortener

COPY --from=builder /url-shortener/binary .

EXPOSE 3000

ENTRYPOINT [ "./binary" ]