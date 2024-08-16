package repository

import (
	"github.com/AndresBC-Dev/parfum/entity"
	"gorm.io/gorm"
)

type PerfumeRepository interface {
	Create(parfum entity.Parfum) error
}

type PerfumeRepositoryImpl struct {
	DB *gorm.DB
}

func NewPerfumeRepository(db *gorm.DB) PerfumeRepository {
	return &PerfumeRepositoryImpl{DB: db}
}

func (p *PerfumeRepositoryImpl) Create(parfum entity.Parfum) error {
	return p.DB.Create(&parfum).Error
}
