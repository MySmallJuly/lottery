package service

import (
	"lottery/dao"
	"lottery/models"
)

type BlackIpService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	UpDate(data *models.LtGift,columns []string) error
	Create(data *models.LtGift) error
}

type blackIpService struct {
	dao *dao.BlackIpDao
}

func NewBlackIpService() BlackIpService {
	return &blackIpService{
		dao: dao.NewBlackIpDao(nil),
	}
}

func (s *blackIpService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}
func (s *blackIpService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *blackIpService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}
func (s *blackIpService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *blackIpService) UpDate(data *models.LtGift,columns []string) error {
	return s.dao.UpDate(data,columns)
}
func (s *blackIpService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}