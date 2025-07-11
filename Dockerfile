FROM golang:1.24.4-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o /app/js-runner-service ./cmd/main.go


FROM alpine:latest

RUN apk add --no-cache ca-certificates docker-cli
WORKDIR /app
COPY --from=builder /app/js-runner-service .
EXPOSE 704

CMD ["./js-runner-service"]
