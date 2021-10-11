package controller

import (
	"fmt"
	"gin/common"
	"gin/model"
	"gin/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	tel := c.PostForm("tel")
	password := c.PostForm("password")

	// 数据验证
	if len(tel) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码过短",
		})
		return
	}
	// 如果名称没有传 给一个随机的字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, tel, password)
	// 判断手机号是否存在
	if isTelExist(DB, tel) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户已经存在",
		})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:     name,
		Tel:      tel,
		Password: password,
	}
	res := DB.Create(&newUser)
	fmt.Println(res)
	// 返回结果
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func isTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
