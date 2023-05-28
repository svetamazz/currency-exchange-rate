FROM golang:1.20.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o server

# Expose the port on which the server will run
EXPOSE 8080

# Run the Go application with environment variables
CMD ["sh", "-c", "source .env && ./server"]
