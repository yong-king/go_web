package main

import (
	"project/dao"
	"project/models"
	"project/routers"
)

func main() {

	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 连接表单
	dao.DB.AutoMigrate(&models.Todo{})

	// 路由
	router := routers.SetUpRouter()

	// 运行服务
	router.Run(":9090")
}
