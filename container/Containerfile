FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY go.* .
COPY . .

#RUN echo "Running Test Suite"
#RUN go test -cover ./...

RUN echo "Building binary"
RUN go build -o bin ./cmd/client/main.go

RUN echo "Exposing Port: 3030"
EXPOSE 3030

RUN echo "Running binary"
ENTRYPOINT [ "./bin", "-prod", "-port", "3030" ]

FROM alpine:latest AS server

WORKDIR /server

RUN echo "Copying files form build stage..."
COPY --from=build /app/bin /server/athaeneum
COPY --from=build /app/public /server/public
COPY --from=build /app/.env /server/.env

RUN echo "Creating logs dir..."
RUN mkdir logs

RUN echo "Executing binary..."
ENTRYPOINT [ "./athaeneum", "-prod" ]