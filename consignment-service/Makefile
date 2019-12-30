build:
	protoc --go_out=plugins=grpc:. -I. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-service-consignment .

run:
	docker run -p 50051:50051 shippy-service-consignment
