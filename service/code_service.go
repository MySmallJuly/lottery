package service

import (
	"lottery/dao"
	"lottery/models"
)

type CodeService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	UpDate(data *models.LtGift,columns []string) error
	Create(data *models.LtGift) error
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewCodeDao(nil),
	}
}

func (s *codeService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}
func (s *codeService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *codeService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}
func (s *codeService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *codeService) UpDate(data *models.LtGift,columns []string) error {
	return s.dao.UpDate(data,columns)
}
func (s *codeService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}