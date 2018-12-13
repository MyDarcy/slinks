package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slinks/models"
)

// 跳转函数，GET
func Redirect(c *gin.Context){
	slink := c.Param("slink")
	//c.String(http.StatusOK, slink)
	has,llink := getLlink(slink)
	if !has{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "get long link fail. Please check again your address carefully.",
			"data": llink,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, llink)
}

//根据短连接后缀拿到可以直接跳转的长连接
func getLlink(slink string)(bool, string){
	var err error
	var has bool
	redisConn := models.RedisConn
	//在这里加上先从redis取的逻辑，没有再落DB
	isInRedis, resRedis := models.GetLinkInRedis("s", slink)
	if isInRedis{
		return true, resRedis
	}
	//redis找不到，则落mysql查找
	engine := models.Orm
	linkMap := new(models.LinkMap)
	linkMap.Slink = slink
	has, err = engine.Get(linkMap)
	if err!=nil{
		models.ErrIntoFile("func getLlink", err.Error())
		return false, err.Error()
	}
	if !has{
		return false, ""
	}
	//如果mysql有，redis没有，可以将slink-llink写到redis, 也可以再设置过期时间
	_,err = redisConn.Do("setnx", "slinks:s:" + linkMap.Slink, linkMap.Llink)
	if err != nil{
		models.ErrIntoFile("func getLlink", err.Error())
	}
	return true,linkMap.Llink
}

