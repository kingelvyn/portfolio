# Use the official Go image to build
FROM golang:1.21 as builder

WORKDIR /app

# Copy go files and static/template files
COPY . .

# Build the Go app
RUN go build -o portfolio .

# Use a minimal image to run it
FROM debian:bullseye-slim

WORKDIR /app

# Copy binary and necessary folders
COPY --from=builder /app/portfolio .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Expose port 3000
EXPOSE 3000

CMD ["./portfolio"]