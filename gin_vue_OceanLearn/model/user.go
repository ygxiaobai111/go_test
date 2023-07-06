package model

type Go_gin_User struct {
	Id        int
	Name      string `xorm:"varchar(20)"`
	Telephone string `xorm:"varchar(11)"`
	Password  string `xorm:"varchar(20)"`
}
