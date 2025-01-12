# Use the official Go image as the build stage
FROM golang:1.23.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY app/go.mod app/go.sum ./

# Download all dependencies (cache dependencies layer)
RUN go mod download

# Copy the source code into the container
COPY . .

# Move to the directory containing the Go application
WORKDIR /app/app

# Build the application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a minimal base image for the final container
FROM alpine:latest

# Install necessary certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder
COPY --from=builder /app/app/main .

# Expose the port your Gin app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
