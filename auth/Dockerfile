# Use the official Golang image as the base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files (go.mod and go.sum) to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go

# Use a lightweight alpine image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the build stage to the final stage
COPY --from=0 /app/app .

# Expose the port on which the application will run
EXPOSE 8000

# Execute the application
CMD ["./app"]