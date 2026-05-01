package handler

import pb "github.com/uber-clone/xxx-service/proto"

type GrpcHandler struct {
    pb.UnimplementedNotificationServiceServer
}

func NewGrpcHandler() *GrpcHandler {
    return &GrpcHandler{}
}
