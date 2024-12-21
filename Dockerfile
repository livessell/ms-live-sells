# Step 1: Build the binary
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . ./

# Copy the .env file into the container
COPY .env .env

# Compile the Go binary with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o main

# Ensure the binary is executable
RUN chmod +x main

# Step 2: Final container using the latest Debian version
FROM debian:bookworm-slim

# Install the necessary dependencies
RUN apt-get update && apt-get install -y libc6

RUN apt-get update && apt-get install -y ca-certificates

# Set the working directory inside the final container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Expose the port the application will listen on
EXPOSE 8080

# Default command to run the Go application
CMD ["./main"]