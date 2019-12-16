package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"

	pb "github.com/ab22/shippy-service-consignment/proto/consignment"
)

const (
	address         = "localhost:50051"
	consignmentFile = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var (
		consignment *pb.Consignment
		data, err   = ioutil.ReadFile(file)
	)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	log.Println("Dialing server...")
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not dial: %v", err)
	}

	log.Println("Parsing consignment information...")
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)
	consignment, err := parseFile(consignmentFile)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	log.Println("Executing creating consignment command...")
	r, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("create consignment failed: %v", err)
	}

	log.Println("Created consignment:", r.Created)
}
