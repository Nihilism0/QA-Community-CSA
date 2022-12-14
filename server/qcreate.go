package server

import (
	"CSAwork/model"
	"CSAwork/pb/proto"
	"context"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net"
	"time"
)

func Qcreate() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50058")
	log.Println("已监听50058,作用是新建问题微服务")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //获取新服务示例
	proto.RegisterCreateQuestionServer(s, &CreateQuestionServer{})
	// 开始处理
	err = s.Serve(lis)
	log.Println("开始处理")
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type CreateQuestionServer struct {
	proto.UnimplementedCreateQuestionServer
}

func (s *CreateQuestionServer) mustEmbedUnimplementedCreateQuestionServer() {
	//TODO implement me
	panic("implement me")
}

func (s *CreateQuestionServer) CreateQuestion(ctx context.Context, req *proto.QcreateReq) (*proto.QcreateResp, error) {
	log.Println("recv:", req.UserName, req.Message)
	db, _ := gorm.Open(mysql.New(mysql.Config{ //配置
		DSN: "csademo:BtBBmGBd7Y8BK7EA@tcp(49.234.42.190:3306)/csademo?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(10) //数据池
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	TestQuestion := model.Question{
		Questioner: req.UserName,
		Message:    req.Message,
	}
	db.Model(&model.Question{}).Create(&TestQuestion)
	resp := &proto.QcreateResp{}
	resp.OK = true
	resp.QuestionID = uint32(TestQuestion.ID)
	return resp, nil
}
