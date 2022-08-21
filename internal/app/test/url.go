package test

import (
	"github.com/gin-gonic/gin"
)

func Url(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("test", func(c *gin.Context) {
			c.String(200, "ok")
		})
		v1.GET("test1", test)
		v1.GET("test2", testSql)
		v1.GET("test3", testRedis)
		v1.GET("login", testLogin)
		v1.GET("logout", testLogout)
	}
}
