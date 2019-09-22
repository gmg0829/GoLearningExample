package router

import (
    "github.com/gin-gonic/gin"
    . "GoLearningExample/src/gin/example/api/apis"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/users", Users)

    router.POST("/user", Store)

    router.PUT("/user/:id", Update)

    router.DELETE("/user/:id", Destroy)

    return router
}