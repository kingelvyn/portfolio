# Use the official Go image to build
FROM golang:1.21-alpine as builder

WORKDIR /app

# Install packages for build stage
RUN apk add --no-cache git

# Copy go files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy rest of code
COPY . .

# Disables CGO
ENV CGO_ENABLED=0
# Build the Go app
RUN go build -o portfolio .

# Use a minimal image to run
FROM alpine:latest

WORKDIR /app

# Install adduser for final image
RUN apk add --no-cache ca-certificates

# Adding user
RUN adduser -D -h /home/elvyn elvyn

# Copy binary & necessary assets
COPY --from=builder /app/portfolio .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/content ./content

# Expose port 3000
EXPOSE 3000

# Run as user
USER elvyn

# Start server
CMD ["./portfolio"]