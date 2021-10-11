package main

import (
	"gin/common"
	"gin/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// test
	db := common.GetDB()
	db.AutoMigrate(&model.User{})
	r := gin.Default()
	r = CollectRoute(r)
	r.Run(":9090")
}
