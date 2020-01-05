package main

import (
	"context"
	"log"
	"os"

	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/ab22/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	port             = ":50051"
	defaultDatastore = "datastore:27017"
)

func getEnvOr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	var (
		uri              = getEnvOr("DB_HOST", defaultDatastore)
		srv              = micro.NewService(micro.Name("shippy.consignment.service"))
		vesselClient     = vesselProto.NewVesselServiceClient("shippy.vessel.service", srv.Client())
		mongoClient, err = CreateMongoClient(uri)
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
