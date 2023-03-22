package default_router

import "github.com/gin-gonic/gin"

func AttachRoutes(router *gin.Engine) {

	router.POST("/ping", ping())

}
