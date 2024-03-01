package server

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/ad", adsCreateApiHandler)
	r.GET("/api/v1/ad", adsSearchApiHandler)

	return r
}

func adsCreateApiHandler(ctx *gin.Context) {

}

func adsSearchApiHandler(ctx *gin.Context) {

}
