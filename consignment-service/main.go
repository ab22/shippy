package main

import (
	"context"
	"log"

	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	"github.com/ab22/shippy/internal/db"
	"github.com/ab22/shippy/internal/env"
	vesselProto "github.com/ab22/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	port             = ":50051"
	defaultDatastore = "datastore:27017"
)

func main() {
	var (
		uri              = env.LookupEnvOr("DB_HOST", defaultDatastore)
		srv              = micro.NewService(micro.Name("shippy.consignment.service"))
		vesselClient     = vesselProto.NewVesselServiceClient("shippy.vessel.service", srv.Client())
		mongoClient, err = db.NewMongoClient(uri)
	)

	if err != nil {
		log.Panic("Could not create mongo client:", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	log.Println("Starting server...")
	srv.Init()

	log.Println("Creating repository...")
	var (
		consignmentCollection = mongoClient.Database("shippy").Collection("consignments")
		repo                  = &MongoRepository{consignmentCollection}
		h                     = &handler{repo, vesselClient}
	)

	log.Println("Registering handler...")
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	log.Println("Running service...")
	if err := srv.Run(); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
