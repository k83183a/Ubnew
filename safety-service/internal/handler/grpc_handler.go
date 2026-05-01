package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedSafetyServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
