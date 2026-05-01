package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedAi-conciergeServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
