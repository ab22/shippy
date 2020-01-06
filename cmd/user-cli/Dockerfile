# Builder
FROM golang:alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOD=linux go build -a -installsuffix csgo -o user-cli ./cmd/user-cli

# App
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/user-cli .
CMD ["./user-cli"]
