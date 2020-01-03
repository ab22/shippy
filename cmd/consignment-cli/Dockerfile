# Builder
FROM golang:apine as builder
RUN apk --no-cache add git
WORKDIR /app/shippy-consignment-cli
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOD=linux go build -a -installsuffix csgo -o shippy-consignment-cli

# App
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app

ADD consignment.json /app/consignment.json
COPY --from=builder /app/shippy-consignment-cli/shippy-consignment-cli .

CMD ["./shippy-consignment-cli"]
