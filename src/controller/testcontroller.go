package controller

import (
	"github.com/gin-gonic/gin"
	"involutiongo/src/cache"
	"reflect"
	"time"
)

// https://learnku.com/docs/gin-gonic/2019/quickstart/6151

func init() {
	r := gin.Default()
	//r.GET("/set", set)
	//r.GET("/get", get)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}

func set(c *gin.Context) {
	value := c.Query("value")
	redisOps := cache.RedisOps{}
	redisOps.Set("test", value, time.Second*1000)
	c.JSON(200, gin.H{
		"status": "200",
		"data":   nil,
	})
}
func get(c *gin.Context) {
	key := c.Query("key")
	redisOps := cache.RedisOps{}
	value := redisOps.Get(key, reflect.String)
	c.JSON(200, gin.H{
		"status": "200",
		"data":   value,
	})
}
