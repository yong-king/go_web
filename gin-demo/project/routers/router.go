package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/contorller"
)

func SetUpRouter() *gin.Engine {
	route := gin.Default()

	// 加载html路径
	route.LoadHTMLFiles("./templates/index.html")
	// 加载静态文件
	route.Static("/css", "./static/css")
	route.Static("/js", "./static/js")
	route.Static("/static/fonts", "./static/fonts")
	// 获取html页面
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 代办事务
	v1Serve := route.Group("/v1")
	{
		// 添加事务
		v1Serve.POST("/todo", contorller.AddTodo)

		// 查询所有事务
		v1Serve.GET("/todo", contorller.GetAllTodo)

		// 查询单个事务
		v1Serve.GET("/todo/:id", contorller.GetTodoByID)

		// 修改事务
		v1Serve.PUT("/todo/:id", contorller.UpdataTodoByID)

		// 删除事务
		v1Serve.DELETE("/todo/:id", contorller.DeleteTodoByID)
	}
	return route
}
