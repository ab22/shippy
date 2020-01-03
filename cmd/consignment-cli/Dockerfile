# Builder
FROM golang:alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOD=linux go build -a -installsuffix csgo -o consignment-cli ./cmd/consignment-cli

# App
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
ADD ./cmd/consignment-cli/consignment.json /app/consignment.json
COPY --from=builder /build/consignment-cli .
CMD ["./consignment-cli"]
