package models

type LtYouhui struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	GiftId     int    `xorm:"not null comment('奖品ID') INT(11)"`
	Code       string `xorm:"not null comment('虚拟卷编码') unique VARCHAR(255)"`
	Uri        string `xorm:"not null comment('uri') VARCHAR(50)"`
	SysCreate  int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated int    `xorm:"not null default 0 comment('修改时间') INT(11)"`
	SysStatus  int    `xorm:"not null default 0 comment('状态') SMALLINT(6)"`
}
