package models

type BlackUser struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	Blacktime  int    `xorm:"not null default 0 comment('黑名单到期时间') INT(11)"`
	Realname   string `xorm:"not null default '' comment('联系人') VARCHAR(50)"`
	Address    string `xorm:"not null default '' comment('地址') VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysUpdated int    `xorm:"not null default 0 comment('修改时间') INT(11)"`
	SysIp      int    `xorm:"not null default 0 comment('用户抽奖ip') INT(11)"`
}
