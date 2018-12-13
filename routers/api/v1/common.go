package v1

import (
	"github.com/gin-gonic/gin"
	"slinks/models"
)

func MyReceiver (c *gin.Context) (bool,string) {
	bytes,err := c.GetRawData()
	if err !=nil{
		models.ErrIntoFile("func Receiver", err.Error())
		return false,""
	}
	data := string(bytes)
	return true,data
}

