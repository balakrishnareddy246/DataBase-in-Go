package main

import

(

   "github.com/gin-gonic/gin"

   "strconv"

   "github.com/hashicorp/consul/api"

)
type User struct {

    Id int64 `db:"id" json:"id"`

    Policy string `db:"policy" json:"policy"`

}
func main() {

  r := gin.Default()

  v1 := r.Group("api/v1")

  {

  v1.GET("/policies", GetPolicies)

  v1.GET("/policies/:id", GetPolicies)

  v1.POST("/policies", PostPolicy)

  v1.PUT("/policies/:id", UpdatePolicy)

  v1.DELETE("/policies/:id", DeletePolicy)

  }

  r.Run(":8080")

}
