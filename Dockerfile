# Use an official Golang runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -mod=vendor -o logistic-api

# Expose the port the application runs on
EXPOSE 8080

# Run the application
CMD ["./logistic-api"]
