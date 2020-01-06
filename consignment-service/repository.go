package main

import (
	"github.com/ab22/shippy/internal/db"

	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	mgo "gopkg.in/mgo.v2"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

// ConsignmentRepository -
type ConsignmentRepository struct {
	dbName         string
	collectionName string
	session        *mgo.Session
}

// NewConsignmentRepository --
func NewConsignmentRepository(uri, dbname, collection string) (*ConsignmentRepository, error) {
	session, err := db.NewMongoClient(uri)

	if err != nil {
		return nil, err
	}

	return &ConsignmentRepository{
		dbName:         dbname,
		collectionName: collection,
		session:        session,
	}, nil
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(repo.dbName).C(repo.collectionName)
}

// Create a new consignment.
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetAll consignments.
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var (
		consignments []*pb.Consignment
		err          = repo.collection().Find(nil).All(&consignments)
	)
	return consignments, err
}

// Close the database connection.
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}
