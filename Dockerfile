FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o kubescale ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/kubescale .
COPY config.yaml .
EXPOSE 8080
EXPOSE 9090
CMD ["./kubescale"]