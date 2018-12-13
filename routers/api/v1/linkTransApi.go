package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"net/http"
	"slinks/models"
	"strings"
)

const digits62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


//查询短链接并返回对应结果. POST
func GetLink(c *gin.Context){
	flag,data := MyReceiver(c)
	if !flag{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "get link fail",
			"data": "",
		})
		return
	}
	slink := gjson.Get(data, "slink").String()
	//c.String(http.StatusOK, "Hello %s", slink)
	slink_slice := strings.Split(slink, `/`)
	slink = slink_slice[len(slink_slice)-1]
	//先从redis从查
	isInRedis, resRedis := models.GetLinkInRedis("s", slink)
	if isInRedis{ //redis中能查到的话
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "get link success",
			"data": resRedis,
		})
		return
	}
	//redis中没有，则需要落mysql查
	engine := models.Orm
	redisConn := models.RedisConn
	linkMap := new(models.LinkMap)
	linkMap.Slink = slink
	_,err := engine.Get(linkMap)
	if err!=nil{
		models.ErrIntoFile("func GetLink", err.Error())
	}
	if linkMap.Llink=="" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "corresponding llink not exist!",
			"data": "",
		})
		return
	}
	//写入redis中缓存
	_,err = redisConn.Do("setnx", linkMap.Slink, linkMap.Llink)
	if err != nil{
		models.ErrIntoFile("func getLlink", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "get link success",
		"data": linkMap,
	})
}


// 生成短链接. POST
func GenSlink(c *gin.Context){
	flag,data := MyReceiver(c)
	if !flag{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "parse post-data error",
			"data": "",
		})
		return
	}
	llink := gjson.Get(data, "llink").String()
	//根据长链接生成短链接，这里也做一层缓存，但是这样会有数据冗余的问题，即长->短和短->长数据就各存一份了
	slink := genLinkHelper(llink)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "gen slink success",
		"data": gin.H{
			"slink": packShortLink(slink),
			"llink": llink,
		},
	})
}

// 算法主要参考：https://stackoverflow.com/questions/742013/how-do-i-create-a-url-shortener
func genLinkHelper(llink string) (string){
	engine:=models.Orm
	redisConn := models.RedisConn
	linkMap := new(models.LinkMap)
	linkMap.Llink = llink
	var s string
	var has bool
	var err error
	var res string
	//1. 先看长链接是否已存在
	//1.1 看redis
	has, res = models.GetLinkInRedis("l", llink)
	if has{
		return res
	}
	//1.2 再看mysql
	has, err = engine.Get(linkMap)
	if err!=nil{
		models.ErrIntoFile("func genLinkHelper", err.Error())
	}
	if !has { //还不存在才插入
		_, err = engine.Insert(linkMap)
		if err!=nil{
			models.ErrIntoFile("func genLinkHelper", err.Error())
		}
		//将linkMap.Id仿照 strconv.FormatInt 从10进制转换为62进制（根据digits62）.
		b := 62
		u := linkMap.Id
		var a [64 + 1]byte // +1 for sign of 64bit value in base 2
		i := len(a)
		for u >= b {
			i--
			q := u / b
			a[i] = digits62[uint(u-q*b)]
			u = q
		}
		// u < base
		i--
		a[i] = digits62[uint(u)]
		s = string(a[i:])
		//将slink更新到db
		linkMap.Slink = s
		_, err = engine.Where("Id = ?", linkMap.Id).Update(linkMap)
		if err!=nil{
			models.ErrIntoFile("func genLinkHelper 2", err.Error())
		}
	} else{
		s = linkMap.Slink
	}
	//再将长链接插入到redis中
	redisConn.Do("setnx", "slinks:l:"+llink, s)
	return s
}


func packShortLink(slink string)(string){
	return fmt.Sprintf("http://s.xawei.me/s/%s", slink)
}