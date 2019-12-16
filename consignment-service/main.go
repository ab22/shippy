package main

import (
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc/reflection"

	// Import protobuf code
	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the user of a datastore of some
// kind. We'll replace this with a real implementation later on.
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new confignment
func (repo *Repository) Create(consigment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()

	updated := append(repo.consignments, consigment)
	repo.consignments = updated

	repo.mu.Unlock()

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
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// Save our consignment
	consignment, err := s.repo.Create(req)

	if err != nil {
		return nil, err
	}

	// Return matching the 'Response' message we created in our
	// protobuf definition.
	return &pb.Response{
		Created:     true,
		Consignment: consignment,
	}, nil
}

// GetConsignments -
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{
		Consigments: consignments,
	}, nil
}

func main() {
	log.Println("Starting server...")
	repo := &Repository{}

	// Set-up our gRPC server.
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register our service with the gRPC server; this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
