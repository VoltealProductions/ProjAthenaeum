# Dockerfile
# named this stage as builder ----------------------
FROM golang:1.22.5-alpine AS build

WORKDIR /app

# Get the go.mod and go.sum files.
# Download dependencies into local cache
COPY go.* .
RUN go mod download

# Copy all files needed for a build.
COPY . .

#Test the go code before build
#RUN echo "Running Test Suite"
#RUN go test -cover ./...

# Build the go binary
RUN go build -o bin ./cmd/client/main.go

# new stage -------------------
FROM alpine:latest

WORKDIR /server

# Copy files from previous build stage.
COPY --from=build /app/bin /server/athaeneum
COPY --from=build /app/public /server/public

# Make logs folder and file.
RUN mkdir ./logs
RUN touch ./logs/system.log

# Execute the Binary
ENTRYPOINT [ "/server/athaeneum", "-prod" ]