package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"slinks/models"
)

func HiRedis(c *gin.Context) {
	redisConn := models.RedisConn
	res, err := redis.String(redisConn.Do("get", "hello"))
	if err != nil{
		models.ErrIntoFile("func HiRedis", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": res,
		"msg": "receiving info from redis.",
	})
}

func Hi (c *gin.Context){
	engine := models.Orm

	linkMap := new(models.LinkMap)
	linkMap.Slink = `s.xawei.me/01s3yy`
	has, err := engine.Get(linkMap)
	if err!=nil{
		models.ErrIntoFile("func Hi", err.Error())
		fmt.Println("func Hi", err.Error())
	}
	if !has{
		linkMap.Llink = "xawei.me"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "hi, " + linkMap.Llink,
		"data": "data here",
	})
}