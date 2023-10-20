# Use the official golang image as the base image
FROM golang:1.21.3 as builder

# Set the working directory in the container
WORKDIR /app

# Copy the rest of the application files to the working directory
COPY . .

# Download the dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 go build -o main .

# Use the light weight alpine image as the base image
FROM alpine:3.18.4

# Set the working directory in the container
WORKDIR /app

# Copy the built application from the builder image
COPY --from=builder /app/main .
COPY --from=builder /app/pkg/email/templates ./pkg/email/templates
COPY --from=builder /app/docs ./docs

# Expose port 3000 to the host
EXPOSE 3000
ENV GIN_MODE=release

# Run the application
CMD ["./main"]
