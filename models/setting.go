package models

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_"github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var (
	Orm *xorm.Engine
	RedisConn redis.Conn
	Cfg *ini.File
	err error
	User string
	Password string
	Host string
	Name string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
)


func InitStarter() {
	initConfig()
	initDB()
}

func initConfig()  {
	Cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/config.ini': %v", err)
	}
	HTTPPort = 9090
	ReadTimeout = 60
	WriteTimeout = 60
}

func initDB(){
	sectionName := "mysql"
	var user = Cfg.Section(sectionName).Key("USER").MustString("root")
	var passwd = Cfg.Section(sectionName).Key("PASSWORD").MustString("123")
	var host = Cfg.Section(sectionName).Key("HOST").MustString("127.0.0.1:3306")
	var name = Cfg.Section(sectionName).Key("NAME").MustString("test")

	Orm,err = xorm.NewEngine("mysql", user+":"+passwd+"@tcp("+host+")/"+name)
	//Dbx, err = sqlx.Open("mysql", user+":"+passwd+"@tcp("+host+")/"+name)
	if err != nil {
		fmt.Println(err)
		fmt.Println(user+":"+passwd+"@tcp("+host+")/"+name)
	}
	//设置mapper
	Orm.SetMapper(core.SameMapper{})

	//初始化redis连接
	sectionName = "redis"
	var ip = Cfg.Section(sectionName).Key("IP").MustString("127.0.0.1")
	var port = Cfg.Section(sectionName).Key("PORT").MustString("6379")
	var passwdRedis = Cfg.Section(sectionName).Key("PASSWORD").MustString("123456")
	redisAddress := ip + ":" + port
	RedisConn, err = redis.Dial("tcp", redisAddress, redis.DialPassword(passwdRedis))
	if err != nil {
		ErrIntoFile("func initDB", err.Error())
	}
}