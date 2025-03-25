package contorller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/models"
)

/*
	控制页面上的增删改查
	URL --> controller --> logic --> model
	请求来了 --> 控制器 --> 业务逻辑 --> 模型层的增删改查
*/

// 添加一个事务
func AddTodo(c *gin.Context) {
	var to models.Todo
	err := c.BindJSON(&to)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err = models.CreatTodo(&to); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to create"})
		return
	}
	c.JSON(http.StatusOK, to)
}

// 获取所有事务
func GetAllTodo(c *gin.Context) {
	// 从数据库中拿到数据
	todo, err := models.FindAllTodo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fail to find"})
		return
	}
	// 将数据传到回页面
	c.JSON(http.StatusOK, todo)
}

// 通过id查询事务
func GetTodoByID(c *gin.Context) {
	// 获取url参数
	id := c.Param("id")

	// 查询记录
	todo, err := models.GetATodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdataTodoByID(c *gin.Context) {
	// 获取url中id
	id := c.Param("id")

	// 从数据库中查询
	err := models.UpdateTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	// 返回更新后到数据
	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}

// 通过id删除事务
func DeleteTodoByID(c *gin.Context) {
	// 获取id
	id := c.Param("id")

	// 删除记录
	err := models.DeleteTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
