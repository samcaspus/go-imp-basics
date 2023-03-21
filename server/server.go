package server

import "github.com/gin-gonic/gin"

func InitServer() *gin.Engine {

	router := gin.Default()
	return router

}
