package models

type LtGift struct {
	Id           int    `xorm:"not null pk autoincr INT(10)"`
	Title        string `xorm:"not null default '' comment('奖品名称') VARCHAR(255)"`
	PrizeNum     int    `xorm:"not null default 0 comment('奖品数量') INT(11)"`
	LeftNum      int    `xorm:"not null default 0 comment('剩余数量') INT(10)"`
	PrizeCode    string `xorm:"not null default '' comment('中奖概率') VARCHAR(50)"`
	PrizeTime    int    `xorm:"not null default 0 comment('发奖日期') INT(11)"`
	Img          string `xorm:"not null default '' comment('奖品图片') VARCHAR(255)"`
	Displayorder int    `xorm:"not null default 0 comment('位置序号') INT(10)"`
	Gtype        int    `xorm:"not null default 0 comment('奖品类型') INT(11)"`
	Gdata        string `xorm:"not null default '' comment('虚拟币数量') VARCHAR(255)"`
	TimeBegin    int    `xorm:"not null default 0 comment('开始时间') INT(11)"`
	TimeEnd      int    `xorm:"not null default 0 comment('结束时间') INT(11)"`
	PrizeData    string `xorm:"not null comment('发奖计划') MEDIUMTEXT"`
	PrizeBegin   int    `xorm:"not null default 0 comment('发奖周期开始') INT(11)"`
	PrizeEnd     string `xorm:"not null default '' comment('发奖周期结束') VARCHAR(1022)"`
	SysStatus    int    `xorm:"not null default 0 comment('状态') SMALLINT(6)"`
	SysCreated   int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	SysUpdated   int    `xorm:"not null default 0 comment('修改时间') INT(11)"`
	SysIp        string `xorm:"not null default '' comment('操作人ip') VARCHAR(50)"`
}
