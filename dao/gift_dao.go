package dao

import (
	"lottery/models"
	"github.com/go-xorm/xorm"
	"log"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine:engine,
	}
}

func (d *GiftDao) Get(id int)  *models.LtGift{
	data := &models.LtGift{Id:id}
	ok, err := d.engine.Get(data)
	if ok &&err ==nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *GiftDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift,0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)
	if err != nil {
		log.Println(err)
	}
	return datalist
}

func (d *GiftDao) CountAll()  int64{
	num, err := d.engine.Count(&models.LtGift{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *GiftDao) Delete(id int) error {
	data := &models.LtGift{Id:id,SysStatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *GiftDao) UpDate(data *models.LtGift,columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *GiftDao) Create(data *models.LtGift) error {
	_,err := d.engine.Insert(data)
	return err
}


















