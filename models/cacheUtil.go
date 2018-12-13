package models

import "github.com/gomodule/redigo/redis"

// ******************* 以下为redis使用的辅助函数 *******************
func GetLinkInRedis(linkType string, slinkOrllink string) (bool, string) {
	redisConn := RedisConn
	var linkKeyPrefix string
	switch linkType {
	case "l":
		linkKeyPrefix = "slinks:l:"
	case "s":
		linkKeyPrefix = "slinks:s:"
	default:
		ErrIntoFile("func GetLinkInRedis", "linkType Only support l or s!")
		return false, "linkType Only support l or s!"
	}
	//自定义命名规范
	resRaw, err := redisConn.Do("get", linkKeyPrefix + slinkOrllink)
	if err!=nil{
		ErrIntoFile("func getLlink", err.Error())
		return false,err.Error()
	}
	if resRaw!=nil{ //如果在redis里面有slink-llink键值对，则返回了
		res,_ := redis.String(resRaw, err)
		return true, res
	} else{
		return false, ""
	}
}

