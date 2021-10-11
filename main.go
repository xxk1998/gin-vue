package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Tel      string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	db.AutoMigrate(&User{})
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {

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
			name = RandomString(10)
		}
		log.Println(name, tel, password)
		// 判断手机号是否存在
		if isTelExist(db, tel) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户已经存在",
			})
			return
		}
		// 创建用户
		newUser := User{
			Name:     name,
			Tel:      tel,
			Password: password,
		}
		res := db.Create(&newUser)
		fmt.Println(res)
		// 返回结果
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.Run(":9090")
}

func RandomString(n int) string {
	var letters = []byte("WEFIHBWEJHEWFBIUfbnskfjbwefbe")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

func InitDB() *gorm.DB {
	dsn := "root:xxxx1111@tcp(127.0.0.1:3306)/db_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB连接失败")
	}
	return db
}

func isTelExist(db *gorm.DB, tel string) bool {
	var user User
	db.Where("tel = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
