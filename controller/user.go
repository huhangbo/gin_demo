package controller

import (
	"github.com/gin-gonic/gin"
	"go_test/common"
	"go_test/model"
	"go_test/util"
	"log"
)

func Register(ctx *gin.Context)  {
	db := common.InitDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	log.Println(name, telephone, password)
	if len(telephone) != 11 {
		ctx.JSON(422, gin.H{
			"code": 422,
			"message": "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(422, gin.H{
			"code": 422,
			"message": "密码长度不能低于6位",
		})
		return
	}
	if util.IsTelephoneExist(db, telephone) {
		ctx.JSON(422, gin.H{
			"code": 422,
			"message": "用户已经存在",
		})
		return
	}
	if len(name) == 0 {
		name = util.RandString(8)
	}
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: password,
	}
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return
	}
	db.Create(&newUser)
	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
}
