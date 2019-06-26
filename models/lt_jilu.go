package models

type LtJilu struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	GiftId     int    `xorm:"not null default 0 comment('奖品ID') INT(11)"`
	GiftName   string `xorm:"not null default '' comment('奖品名称') VARCHAR(255)"`
	GiftType   int    `xorm:"not null default 0 comment('奖品类型') INT(10)"`
	Uid        int    `xorm:"not null default 0 comment('用户id') INT(11)"`
	Username   string `xorm:"not null default '' comment('用户名') VARCHAR(50)"`
	PrizeCode  int    `xorm:"not null default 0 comment('抽奖编号') INT(11)"`
	GiftData   int    `xorm:"not null default 0 comment('获奖信息') INT(10)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysIp      int    `xorm:"not null default 0 comment('用户抽奖ip') INT(11)"`
	SysStatus  int    `xorm:"not null default 0 comment('状态') SMALLINT(6)"`
}
