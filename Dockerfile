# Use the official golang image as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application files to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose port 3000 to the host
EXPOSE 3000

# Run the application
CMD ["./main"]