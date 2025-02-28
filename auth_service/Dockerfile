# Example Dockerfile for auth_service

FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth_service ./cmd

FROM scratch
COPY --from=builder /app/auth_service /auth_service
EXPOSE 8081
ENTRYPOINT ["/auth_service"]
