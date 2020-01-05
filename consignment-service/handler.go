package main

import (
	"context"

	pb "github.com/ab22/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/ab22/shippy/vessel-service/proto/vessel"
)

type handler struct {
	repo         repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := h.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}

	if err = h.repo.Create(req); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	req.VesselId = vesselResponse.Vessel.Id
	return nil
}

// GetConsignments -
func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) (err error) {
	res.Consigments, err = h.repo.GetAll()
	return err
}
