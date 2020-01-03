# Build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/ab22/shippy/consignment-service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-consignment-service .

# Production stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder go/src/github.com/ab22/shippy/consignment-service/shippy-consignment-service .
CMD ["./shippy-consignment-service"]
