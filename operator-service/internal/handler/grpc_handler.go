package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedOperatorServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
