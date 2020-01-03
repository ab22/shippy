# Builder Stage
FROM golang:alpine as builder
RUN apk --no-cache add git
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-vessel-service .

# Production Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-vessel-service .
CMD ["./shippy-vessel-service"]
