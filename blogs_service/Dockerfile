# Example Dockerfile for blogs_service

FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o blog_service ./cmd

FROM scratch
COPY --from=builder /app/blog_service /blog_service
EXPOSE 8082
ENTRYPOINT ["/blog_service"]
