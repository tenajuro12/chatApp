# gateway/Dockerfile
FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod  ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gateway ./

FROM scratch
COPY --from=builder /app/gateway /gateway
EXPOSE 8080
ENTRYPOINT ["/gateway"]
