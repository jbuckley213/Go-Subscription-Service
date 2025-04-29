FROM golang:1.23.3

WORKDIR /app

# Copy the Go module files
COPY go.mod . 

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp .

# Expose the port your application listens on (if any)
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
