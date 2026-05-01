package service

import "github.com/uber-clone/appointment-service/internal/repository"

type Service struct {
    repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
    return &Service{repo: repo}
}
