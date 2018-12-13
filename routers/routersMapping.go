package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slinks/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/index.html")
	})
	api := r.Group("/s")
	{
		api.GET(":slink",v1.Redirect)
	}

	apiv1 := r.Group("/v1")
	{
		apiv1.GET("hi", v1.Hi)
		apiv1.GET("hiredis", v1.HiRedis)
		apiv1.POST("getlink", v1.GetLink)
		apiv1.POST("genslink", v1.GenSlink)
	}
	return r
}