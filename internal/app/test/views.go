package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-jy-sso-test/internal/pkg/apollo"
	"go-jy-sso-test/internal/pkg/mysql"
	"go-jy-sso-test/internal/pkg/rdssentinels"
	"go-jy-sso-test/internal/pkg/requests"
	"time"
)

// @Tags Test API
// @Summary List apollo some config
// @Description get apollo config
// @Accept  json
// @Produce  json
// @Success 200 {array} Response
// @Header 200 {string} Response
// @Failure 400,404 {object} string "Bad Request"
// @Router /v1/test1 [get]
func test(c *gin.Context) {
	c.JSON(200, Response{
		Code: 200,
		Data: map[string]string{
			"RedisSentinels": apollo.Config.RedisSentinelAddress,
			"RedisCluster":   apollo.Config.RedisMasterName,
		},
		Msg: "success",
	})
}

func testSql(c *gin.Context) {
	ticketData := make([]RobotTicket, 0)

	if err := mysql.Engine.Find(&ticketData); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"ticket_data": ticketData,
	})
}

func testRedis(c *gin.Context) {
	//rds := rdssentinels.NewRedis(nil)
	//rds.SetKey("testKey", "this is test demo", 3600 * time.Second)
	//result := rds.GetKey("testKey")

	rdssentinels.RedisConfig.SetKey("testKey", "this is test demo", 3600*time.Second)
	result := rdssentinels.RedisConfig.GetKey("testKey")

	c.JSON(200, gin.H{
		"redis_testKey": result.Val(),
	})
}

func testLogout(c *gin.Context) {
	url := fmt.Sprintf("http://go-test.xxx.xxx/v1/login")
	redirectUrl := fmt.Sprintf("http://sso-api.xxx.xxx/user/feishu/logout?redirect_uri=%s&app_name=test1", url)
	c.Redirect(302, redirectUrl)
}

func testLogin(c *gin.Context) {
	jyToken, err := c.Cookie("xxx_token")
	if err != nil {
		url := fmt.Sprintf("http://go-test.xxx.xxx/v1/login")
		redirectUrl := fmt.Sprintf("http://sso-api.xxx.xxx/user/feishu/login?redirect_uri=%s&app_name=test1", url)
		c.Redirect(302, redirectUrl)
	} else {
		var result SsoData
		checkTokenUrl := fmt.Sprintf("http://sso-api.xxx.xxx/user/token?token=%s", jyToken)
		data := map[string]interface{}{}
		headers := map[string]string{"Content-Type": "application/json"}

		err := requests.RequestMethod(checkTokenUrl, "GET", headers, data, &result)
		if err != nil {
			fmt.Println(err)
		} else {
			if result.Code == 200 {
				c.JSON(200, gin.H{
					"msg":  "this is go test",
					"data": result,
				})
			} else {
				c.Redirect(302, "redirectLogin")
			}
		}
	}
}
