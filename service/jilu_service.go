package service

import (
	"lottery/dao"
	"lottery/models"
)

type JiLuService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	UpDate(data *models.LtGift,columns []string) error
	Create(data *models.LtGift) error
}

type jiLuService struct {
	dao *dao.JiLuDao
}

func NewJiLuService() JiLuService {
	return &jiLuService{
		dao: dao.NewJiLuDao(nil),
	}
}

func (s *jiLuService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}
func (s *jiLuService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *jiLuService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}
func (s *jiLuService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *jiLuService) UpDate(data *models.LtGift,columns []string) error {
	return s.dao.UpDate(data,columns)
}
func (s *jiLuService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}