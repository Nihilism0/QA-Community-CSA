package api

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/pb/proto"
	"CSAwork/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func Qcreate(c *gin.Context) {
	if err := c.ShouldBind(&model.Question{}); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	message := c.PostForm("message")
	conn, err := grpc.Dial("localhost:50058", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	pc := proto.NewCreateQuestionClient(conn)
	QcreateResp, _ := pc.CreateQuestion(context.Background(), &proto.QcreateReq{
		UserName: username.(string),
		Message:  message,
	})
	if !QcreateResp.OK {
		utils.RespFail(c, "创建问题失败")
		return
	}

	err = global.RedisDb.SAdd("questionids", strconv.Itoa(int(QcreateResp.QuestionID))).Err()
	if err != nil {
		log.Fatal("add error", zap.Error(err))
	}
	utils.RespSuccess(c, "亲爱的"+username.(string)+",您成功bb了一条==>"+message)
}

func Acreate(c *gin.Context) {
	if err := c.ShouldBind(&model.Answer{}); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	message := c.PostForm("message")
	wow := c.PostForm("question_id")
	questionID, _ := strconv.Atoi(wow)
	conn, err := grpc.Dial("localhost:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	pc := proto.NewCreateAnswerClient(conn)
	AcreateResp, _ := pc.CreateAnswer(context.Background(), &proto.AcreateReq{
		UserName:   username.(string),
		QuestionID: uint32(questionID),
		Message:    message,
	})
	if !AcreateResp.OK {
		utils.RespFail(c, "创建回答失败")
		return
	}
	utils.RespSuccess(c, "亲爱的"+username.(string)+",您成功回答了问题,(但在问题的正确性方面不一定成功)==>"+message)
}

func Qsubmited(c *gin.Context) {
	username, _ := c.Get("username")
	questions := dao.FindQuestionSubmited(username.(string))
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": questions,
	})
}

func Asubmited(c *gin.Context) {
	username, _ := c.Get("username")
	answers := dao.FindAnswerSubmited(username.(string))
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": answers,
	})
}

func Qmodify(c *gin.Context) {
	if err := c.ShouldBind(&model.QuestionModify{}); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	question := c.PostForm("message")
	wow := c.PostForm("id")
	questionid, _ := strconv.Atoi(wow)
	OK := dao.JudgeQuestion(username.(string), questionid)
	if !OK {
		utils.RespFail(c, "这个问题不属于你")
		return
	}
	dao.QuestionModify(question, questionid)
	utils.RespSuccess(c, "修改问题成功!!!")
}

func Amodify(c *gin.Context) {
	if err := c.ShouldBind(&model.AnswerModify{}); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	answer := c.PostForm("message")
	wow := c.PostForm("id")
	answerid, _ := strconv.Atoi(wow)
	OK := dao.JudgeAnswer(username.(string), answerid)
	if !OK {
		utils.RespFail(c, "这个回答不属于你")
		return
	}
	dao.AnswerModify(answer, answerid)
	utils.RespSuccess(c, "修改回答成功!!!")
}
func Qdelete(c *gin.Context) {
	if err := c.ShouldBind(&model.Delete{}); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	wow := c.PostForm("id")
	questionid, _ := strconv.Atoi(wow)
	OK := dao.JudgeQuestion(username.(string), questionid)
	if !OK {
		utils.RespFail(c, "这个问题不属于你")
		return
	}
	dao.QuestionDelete(questionid)
	utils.RespSuccess(c, "删除问题成功!!!")
}

func Adelete(c *gin.Context) {
	if err := c.ShouldBind(&model.Delete{}); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	wow := c.PostForm("id")
	answerid, _ := strconv.Atoi(wow)
	OK := dao.JudgeAnswer(username.(string), answerid)
	if !OK {
		utils.RespFail(c, "这个回答不属于你")
		return
	}
	dao.AnswerDelete(answerid)
	utils.RespSuccess(c, "删除回答成功!!!")
}
func Acomment(c *gin.Context) {
	if err := c.ShouldBind(&model.Comment{}); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	message := c.PostForm("message")
	wow := c.PostForm("id")
	answerid, _ := strconv.Atoi(wow)
	TestComment := model.Comment{
		Model:     gorm.Model{},
		Commenter: username.(string),
		Message:   message,
		Answer_id: uint(answerid),
	}
	global.GlobalDb1.Model(&model.Comment{}).Create(&TestComment)
	utils.RespSuccess(c, "亲爱的"+username.(string)+"用户,您的评论已发送==>"+message)
}
