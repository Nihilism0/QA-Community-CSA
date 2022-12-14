package api

import (
	"CSAwork/dao"
	"CSAwork/model"
	"CSAwork/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Praise(c *gin.Context) {
	if err := c.ShouldBind(&model.Delete{}); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	wow := c.PostForm("id")
	id2, _ := strconv.Atoi(wow)
	id := strconv.Itoa(int(id2))
	flag1 := dao.SelectQuestion(id)
	if !flag1 {
		utils.RespFail(c, "Question is not exist")
		return
	}
	// 验证用户是否点赞
	flag2 := dao.SelectPraiseuser(id, username.(string))
	// 用户点过赞则退出
	if flag2 {
		utils.RespFail(c, "您已点赞,无需重复点赞!")
		return
	}
	dao.Praiseadd(id, username.(string))
	utils.RespSuccess(c, "点赞成功!")
}

func CancelPraise(c *gin.Context) {
	if err := c.ShouldBind(&model.Delete{}); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	wow := c.PostForm("id")
	id2, _ := strconv.Atoi(wow)
	id := strconv.Itoa(int(id2))
	flag1 := dao.SelectQuestion(id)
	if !flag1 {
		utils.RespFail(c, "问题不存在")
		return
	}
	// 验证用户是否点赞
	flag2 := dao.SelectPraiseuser(id, username.(string))
	// 用户没点过赞则退出
	if !flag2 {
		utils.RespFail(c, "你都没点赞呢,想点踩是吧")
		return
	}
	dao.CancelPraise(id, username.(string))
	utils.RespSuccess(c, "取消点赞成功!")
}

func SeePraise(c *gin.Context) {
	if err := c.ShouldBind(&model.Delete{}); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	wow := c.PostForm("id")
	id2, _ := strconv.Atoi(wow)
	id := strconv.Itoa(int(id2))
	number := dao.SeeQpraise(id)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"PraiseSum": number,
	})
}
