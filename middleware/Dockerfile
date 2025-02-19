FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o dns-validator

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/dns-validator .
EXPOSE 8080
CMD ["./dns-validator"] 