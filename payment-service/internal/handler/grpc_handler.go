package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedPaymentServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
