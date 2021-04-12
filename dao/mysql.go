package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/anta?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
