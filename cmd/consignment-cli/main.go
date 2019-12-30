package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
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
	log.Println("Creating micro service...")
	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	client := pb.NewShippingServiceClient("shippy.consignment.service", service.Client())

	log.Println("Parsing consignment information...")
	consignment, err := parseFile(consignmentFile)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	log.Println("Executing CreateConsignment command...")
	r, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("create consignment failed: %v", err)
	}

	log.Println("Created consignment:", r.Created)
	log.Println("Executing GetConsignments command...")
	r, err = client.GetConsignments(context.Background(), &pb.GetRequest{})

	if err != nil {
		log.Fatalf("Could not get all consignments: %v", err)
	}

	log.Println("Consignments:")
	for _, v := range r.Consigments {
		log.Println(v)
	}
}
