
    # Start from the official Golang image
    FROM golang:1.24.3-alpine

    # Set the working directory inside the container
    WORKDIR /app

    # Copy go.mod and go.sum files
    COPY go.mod go.sum ./

    # Download dependencies
    RUN go mod download

    # Copy the rest of the application code
    COPY . .

    # Build the Go app
    RUN go build -o server .

    # Expose the ports your app uses
    EXPOSE 3000
    EXPOSE 5000

    # Command to run the executable
   CMD ["/bin/sh", "-c", "go build -o server . && ./server"]
