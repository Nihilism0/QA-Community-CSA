package api

import (
	"CSAwork/dao"
	"CSAwork/model"
	"CSAwork/pb/proto"
	"CSAwork/utils"
	"CSAwork/utils/middleware"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名是否重复
	flag := dao.SelectUser(username)
	fmt.Println(flag)
	if flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user already exist~")
		return
	}
	conn, err := grpc.Dial("localhost:50057", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	pc := proto.NewRegisterClient(conn)
	registerResp, _ := pc.Register(context.Background(), &proto.UserReq{
		UserName: username,
		PassWord: password,
	})
	if !registerResp.OK {
		utils.RespFail(c, "add user failed")
	}
	// 以 JSON 格式返回信息
	utils.RespSuccess(c, "add user successful")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}
	//建立链接
	conn, err := grpc.Dial("localhost:50056", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	pc := proto.NewLoginClient(conn)
	loginResp, _ := pc.Login(context.Background(), &proto.UserReq{
		UserName: username,
		PassWord: password,
	})
	if !loginResp.OK {
		utils.RespFail(c, "wrong password")
		return
	}
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Joker",                              // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{"token": tokenString},
	})
}
