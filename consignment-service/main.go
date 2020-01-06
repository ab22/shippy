package main

import (
	"log"

	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	"github.com/ab22/shippy/internal/env"
	vesselProto "github.com/ab22/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	port              = ":50051"
	defaultDatastore  = "datastore:27017"
	defaultDbName     = "shippy"
	defaultCollection = "consignments"
)

func main() {
	log.Println("Starting server...")
	log.Println("Creating repository...")
	var (
		uri        = env.LookupEnvOr("DB_HOST", defaultDatastore)
		dbName     = env.LookupEnvOr("DB_NAME", defaultDbName)
		collection = env.LookupEnvOr("DB_COLLECTION", defaultCollection)
		repo, err  = NewConsignmentRepository(uri, dbName, collection)
	)

	if err != nil {
		log.Panicln("Failed to create repository:", err)
	}
	defer repo.Close()

	log.Println("Creating services...")
	var (
		srv          = micro.NewService(micro.Name("shippy.consignment.service"))
		vesselClient = vesselProto.NewVesselServiceClient("shippy.vessel.service", srv.Client())
		h            = &handler{repo, vesselClient}
	)
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	log.Println("Running service...")
	if err := srv.Run(); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
