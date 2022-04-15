package controller

import (
	"HLRJ/gin_learn/common"
	"HLRJ/gin_learn/model"
	"HLRJ/gin_learn/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	//name := ctx.PostForm("name")
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")
	// gin的bind获取参数
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		fmt.Print("获取失败")
	}
	// 获取参数 名称、手机号和密码
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	fmt.Println(name, telephone, password)
	// 数据验证
	// map[string]interface{} 就是  gin.h
	if len(telephone) != 11 {
		//http常量或者422
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(422, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
		return
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if utils.IsTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)
	//返回结果
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
