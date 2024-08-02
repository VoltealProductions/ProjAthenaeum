FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY go.* .
COPY . .

#RUN echo "Running Test Suite"
#RUN go test -cover ./...

RUN echo "Building binary"
RUN go build -o bin ./cmd/client/main.go

FROM alpine:latest AS server

WORKDIR /server

RUN echo "Copying files form build stage..."
COPY --from=build /app/bin /server/athaeneum
COPY --from=build /app/public /server/public

RUN echo "Creating logs dir..."
RUN mkdir logs

RUN echo "Executing binary..."
ENTRYPOINT [ "/server/athaeneum", "-prod" ]