# We use the official golang image, which contains all the
# correct build tools and libraries. Notice the 'as builder',
# this gives this container a name that we can reference later on.
FROM golang:1.13.5 AS builder
WORKDIR /go/src/github.com/ab22/shippy/consignment-service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-consignment-service .
# Here we're using a second FROM statement, which is strange, but
# this tells Docker to start a new build process with this image.
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder go/src/github.com/ab22/shippy/consignment-service/shippy-consignment-service .
CMD ["./shippy-consignment-service"]
