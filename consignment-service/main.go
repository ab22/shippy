package main

import (
	"context"
	"log"

	// Import protobuf code
	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the user of a datastore of some
// kind. We'll replace this with a real implementation later on.
type Repository struct {
	consignments []*pb.Consignment
}

// Create a new confignment
func (repo *Repository) Create(consigment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consigment)
	repo.consignments = updated
	return consigment, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures
// to give you a better idea.
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Save our consignment
	consignment, err := s.repo.Create(req)

	if err != nil {
		return err
	}

	// Return matching the 'Response' message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments -
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res.Consigments = s.repo.GetAll()
	return nil
}

func main() {
	log.Println("Starting server...")
	repo := &Repository{}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.consignment.service"),
	)

	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	// Run server.
	if err := srv.Run(); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
