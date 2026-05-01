package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedB2bServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
