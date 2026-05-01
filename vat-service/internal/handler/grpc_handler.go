package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedVatServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
