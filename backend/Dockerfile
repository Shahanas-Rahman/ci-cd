# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8081
CMD ["./main"]