package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-jy-sso-test/internal/app/middleware/logger"
	"go-jy-sso-test/internal/app/test"
	"go-jy-sso-test/internal/pkg/apollo"
	"go-jy-sso-test/internal/pkg/mysql"
	"go-jy-sso-test/internal/pkg/snowflake"
)

/*
TODO
1、Event跨实例间通讯
*/

// 初始化 apollo config
func initApolloConfig() {
	var err error
	err = apollo.ReadRemoteConfig()
	if nil != err {
		log.Fatal(err)
	}
}

// 初始化mysql
func initMysql() {
	err := mysql.NewMysql()
	if nil != err {
		log.Fatal(err)
	}
}

// 初始化redis
/*func initRedis() {
	rdssentinels.NewRedis(nil)
}*/

// 初始化log配置
func (s *server) initLog() *gin.Engine {
	logs := logger.LogMiddleware()
	s.App.Use(logs)
	return s.App
}

// 初始化雪花算法
func initSnowFlake() {
	snowflake.InitSnowWorker(1, 1)
}

// 加载gin 路由配置
func (s *server) InitRouter() *gin.Engine {
	test.Url(s.App)
	return s.App
}

// InitSwagger init swagger
func (s *server) InitSwagger() *gin.Engine {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	s.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return s.App
}
