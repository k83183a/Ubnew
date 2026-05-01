package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedFeatureflagServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
