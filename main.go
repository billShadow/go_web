package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"bubble/dao"
	"bubble/models"
	"bubble/routes"
)

func main() {

	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Toto{})

	//注册路由
	r := routes.InitRoute()

	r.Run(":9008")
}
