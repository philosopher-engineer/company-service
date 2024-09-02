# Use an official Golang image to build the application
FROM golang:1.23-bookworm as build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
#RUN go build -o ./main ./cmd/main.g
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/company-api /app/cmd/main.go

# Set environment variables
ENV GIN_MODE=release

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["/app/company-api"]

