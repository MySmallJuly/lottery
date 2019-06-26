package models

type BlackIp struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Ip         string `xorm:"not null default '' comment('用户名') unique VARCHAR(50)"`
	Blacktime  int    `xorm:"not null default 0 comment('ip地址') INT(11)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysUpdated int    `xorm:"not null default 0 comment('修改时间') INT(11)"`
	SysIp      int    `xorm:"not null default 0 comment('用户抽奖ip') INT(11)"`
}
