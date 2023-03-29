package service

import "challenge07/repository"

type Service struct {
	repo repository.RepositoryInterface
}

type ServiceInterface interface {
	BookService
}

func NewService(repo repository.RepositoryInterface) *Service {
	return &Service{repo: repo}
}
