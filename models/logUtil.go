package models

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func LogIntoFile(prefix string, data string){
	fdTime:=time.Now().Format("2006-01-02 15:04:05")
	logFileName := "dailyLog-" + strings.Fields(fdTime)[0] + ".log"
	RealLogFileName := "logs/" + logFileName
	f, err := os.OpenFile(RealLogFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil{
		fmt.Println(err, "in LogIntoFile")
	} else{
		logger := log.New(f, "[" + prefix + "]", log.LstdFlags)
		logger.Println("  |  ",data)
	}
}

func ErrIntoFile(prefix string, data string){
	fdTime:=time.Now().Format("2006-01-02 15:04:05")
	logFileName := strings.Fields(fdTime)[0] + "Err" + ".log"
	RealLogFileName := "logs/" + logFileName
	f, err := os.OpenFile(RealLogFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil{
		fmt.Println(err, "in ErrIntoFile")
	} else{
		logger := log.New(f, "[" + prefix + "]", log.LstdFlags)
		logger.Println(data)
	}
}