# Use the official Go image to build
FROM golang:1.21-alpine as builder

WORKDIR /app

# Copy go files and static/template files
COPY . .

# Disables CGO
RUN CGO_ENABLED=0
# Build the Go app
RUN go build -o portfolio .

# Use a minimal image to run it
FROM alpine:latest

WORKDIR /app

# Copy binary and necessary folders
COPY --from=builder /app/portfolio .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Expose port 3000
EXPOSE 3000

CMD ["./portfolio"]