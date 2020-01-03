build:
	protoc --go_out=plugins=micro:. -I. proto/consignment/consignment.proto
	docker build -t shippy-consignment-service .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		shippy-consignment-service
