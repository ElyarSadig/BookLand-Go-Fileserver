# Use the official Golang image for your base image
FROM golang:latest

WORKDIR /app

# Copy necessary files and folders to the container
COPY . /app

# Build the Go application
RUN go build -o /app/main ./cmd/main.go

# Expose the port your application runs on
EXPOSE 8080

# Define the command to run your application
CMD ["/app/main"]
