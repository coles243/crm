FROM golang:1.23.3

# Set destination for COPY
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o crm-api main.go

# Expose the port the application runs on
EXPOSE 3000

# Run the application
CMD ["./crm-api"]