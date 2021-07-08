package main

import (
	// 导入session包
	"github.com/gin-contrib/sessions"
	"net/http"

	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
	// 导入gin框架包
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)
		// 通过session.Get读取session值
		// session是键值对格式数据，因此需要通过key查询数据
		if session.Get("hello") != "world" {
			// 设置session数据
			session.Set("hello", "world")
			// 删除session数据
			//session.Delete("hello")
			//清空session
			//session.Clear()
			// 保存session数据
			session.Save()
			// 删除整个session
			// session.Clear()
		}
		//转发到另一个url
		c.Request.URL.Path = "/hi"
		r.HandleContext(c)
	})
	r.GET("/hi", func(context *gin.Context) {
		// 初始化session对象
		session := sessions.Default(context)
		context.JSON(http.StatusOK,gin.H{
			"session-hi":session.Get("hello"),
		})
	})

	r.Run(":8000")
}