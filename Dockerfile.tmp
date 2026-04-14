# Use the official Go image to build
FROM golang:1.21-alpine AS builder

ARG CACHE_BUSTER=1

WORKDIR /app

# Install packages for build stage
RUN apk add --no-cache git

# Copy go files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Force Docker to invalidate the cache for everything below
RUN echo "CACHE_BUSTER=${CACHE_BUSTER}"

# Copy rest of code
COPY . .

# Disable CGO and build the app
ENV CGO_ENABLED=0
RUN go build -o portfolio .

# Use a minimal image to run
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies and create user
RUN apk add --no-cache ca-certificates \
    && adduser -D -h /home/elvyn elvyn

# Copy binary and necessary assets
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