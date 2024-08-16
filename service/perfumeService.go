package service

import (
	"github.com/AndresBC-Dev/parfum/entity"
	"github.com/AndresBC-Dev/parfum/repository"
)

type PerfumeService interface {
	Create(parfum entity.Parfum) error
}

type PerfumeServiceImpl struct {
	ParfumRepository repository.PerfumeRepository
}

func NewPerfumeService(parfumRepository repository.PerfumeRepository) PerfumeService {
	return &PerfumeServiceImpl{ParfumRepository: parfumRepository}
}

func (p *PerfumeServiceImpl) Create(parfum entity.Parfum) error {
	return p.ParfumRepository.Create(parfum)
}
