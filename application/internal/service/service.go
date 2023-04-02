package service

import (
	"otus/internal/service/adapters/repository"

	"gitlab.com/kanya384/gotools/logger"
)

type service struct {
	repository repository.Repository
	logger     logger.Interface
	options    Options
}

type Options struct {
}

func (s *service) SetOptions(options Options) {
	if s.options != options {
		s.options = options
	}
}

func New(repository repository.Repository, logger logger.Interface, options Options) (*service, error) {
	service := &service{
		repository: repository,
		logger:     logger,
	}

	service.SetOptions(options)
	return service, nil
}
