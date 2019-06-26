package service

import (
	"lottery/dao"
	"lottery/models"
)

type UserIpService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	UpDate(data *models.LtGift,columns []string) error
	Create(data *models.LtGift) error
}

type userIpService struct {
	dao *dao.UserIpDao
}

func NewUserIpService() UserIpService {
	return &userIpService{
		dao: dao.NewUserIpDao(nil),
	}
}

func (s *userIpService) GetAll() []models.LtGift {
	return s.dao.GetAll()
}
func (s *userIpService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *userIpService) Get(id int) *models.LtGift {
	return s.dao.Get(id)
}
func (s *userIpService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *userIpService) UpDate(data *models.LtGift,columns []string) error {
	return s.dao.UpDate(data,columns)
}
func (s *userIpService) Create(data *models.LtGift) error {
	return s.dao.Create(data)
}