package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"lottery/conf"
	"sync"
	_"github.com/go-sql-driver/mysql"
)
var dbLock sync.Mutex
var masterInstance *xorm.Engine

func INstanceMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMastr()
}

func NewDbMastr() *xorm.Engine {
	sourcename := 	fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.DataBase)
	instance, err := xorm.NewEngine(conf.DriverNama, sourcename)
	if err != nil {
		log.Fatal("dnhelper.NewDbMastr err",err)
		return nil
	}
	instance.ShowSQL(true)

	masterInstance = instance
	return instance

}
