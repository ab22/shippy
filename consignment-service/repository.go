package main

import (
	"context"

	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRepository - Dummy MongoRepository, this simulates the user of a datastore of some
// kind. We'll replace this with a real implementation later on.
type MongoRepository struct {
	collection *mongo.Collection
}

// Create a new confignment
func (repo *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repo.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll consignments
func (repo *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	var (
		consignments []*pb.Consignment
		consignment  *pb.Consignment
		cur, err     = repo.collection.Find(context.Background(), nil, nil)
	)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}

		consignments = append(consignments, consignment)
	}

	return consignments, nil
}
