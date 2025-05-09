# Use the official Go image to build
FROM golang:1.21-alpine as builder

# Install shadow package including adduser
RUN apk add --no-cache shadow

WORKDIR /app

# Copy go files
COPY go.mod go.sum./
RUN go mod download

# Copy rest of code
COPY . .

# Disables CGO
RUN CGO_ENABLED=0
# Build the Go app
RUN go build -o portfolio .

# Use a minimal image to run
FROM alpine:latest

# Adding user
RUN useradd -D -h /home/elvyn elvyn

# Copy binary & necessary assets
COPY --from=builder /app/portfolio .
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

# Expose port 3000
EXPOSE 3000

# Run as user
USER elvyn

# Start server
CMD ["./portfolio"]