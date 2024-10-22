
FROM golang:1.22-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o backend-service ./cmd/main.go

# Run tests
RUN go test -v ./...

FROM alpine:latest
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/backend-service .

# Copy the config.json file from the builder stage
COPY --from=builder /app/config.json .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./backend-service"]
