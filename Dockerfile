FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Use a minimal image for final container
FROM alpine:latest
WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/main .
# Copy .env file (optional)
# COPY --from=builder /app/.env .

# Expose port
EXPOSE 3000

# Run the binary
CMD ["./main"]