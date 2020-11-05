package main

import (
       "github.com/gin-gonic/gin"
)
func main() {
router := gin.Default()
v1 := router.Group("/api/v1/todos")
 {
  v1.POST("/", create)
  v1.GET("/", fetchAll)
  v1.GET("/:id", fetchSingle)
  v1.PUT("/:id", update)
  v1.DELETE("/:id", delete)
 }
 router.Run()
}
