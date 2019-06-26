package dao

import (
	"lottery/models"
	"github.com/go-xorm/xorm"
	"log"
)

type UserIpDao struct {
	engine *xorm.Engine
}

func NewUserIpDao(engine *xorm.Engine) *UserIpDao {
	return &UserIpDao{
		engine:engine,
	}
}

func (d *UserIpDao) Get(id int)  *models.LtGift{
	data := &models.LtGift{Id:id}
	ok, err := d.engine.Get(data)
	if ok &&err ==nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *UserIpDao) GetAll() []models.LtGift {
	datalist := make([]models.LtGift,0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)
	if err != nil {
		log.Println(err)
	}
	return datalist
}

func (d *UserIpDao) CountAll()  int64{
	num, err := d.engine.Count(&models.LtGift{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *UserIpDao) Delete(id int) error {
	data := &models.LtGift{Id:id,SysStatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *UserIpDao) UpDate(data *models.LtGift,columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserIpDao) Create(data *models.LtGift) error {
	_,err := d.engine.Insert(data)
	return err
}

func (d *UserIpDao) GetByIp(ip string) *models.BlackIp {
	datalist := make([]models.BlackIp,0)
	err := d.engine.Where("ip=?",ip).Desc("ip").Limit(1).Find(&datalist)

	if err != nil || len(datalist) < 1 {
		return nil
	} else {
		return &datalist[0]
	}
}















