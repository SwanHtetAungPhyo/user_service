# Use Golang base image
FROM golang:1.23.2-alpine

# Install curl to be able to fetch dependencies
RUN apk update && apk add --no-cache curl

# Create app directory
WORKDIR /app

# Copy Go application code into container
COPY . .


# Install the Go dependencies
RUN go mod tidy
RUN go build -o user-service .

# Use sh instead of bash
ENTRYPOINT ["./user-service"]
