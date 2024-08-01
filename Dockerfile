FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.* .
COPY . .

RUN go build -o bin ./cmd/webserver/main.go

EXPOSE 8080

ENTRYPOINT [ "./bin", "-prod", "-port", "3030" ]