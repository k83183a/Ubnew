package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedDriverServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
