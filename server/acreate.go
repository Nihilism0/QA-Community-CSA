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

func Acreate() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50059")
	log.Println("已监听50059,作用是新建回答微服务")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //获取新服务示例
	proto.RegisterCreateAnswerServer(s, &CreateAnswerServer{})
	// 开始处理
	err = s.Serve(lis)
	log.Println("开始处理")
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type CreateAnswerServer struct {
	proto.UnimplementedCreateAnswerServer
}

func (s *CreateAnswerServer) mustEmbedUnimplementedCreateAnswerServer() {
	//TODO implement me
	panic("implement me")
}

func (s *CreateAnswerServer) CreateAnswer(ctx context.Context, req *proto.AcreateReq) (*proto.AcreateResp, error) {
	log.Println("recv:", req.UserName, req.QuestionID, req.Message)
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
	TestAnswer := model.Answer{
		Model:       gorm.Model{},
		Answerer:    req.UserName,
		Message:     req.Message,
		Question_id: uint(req.QuestionID),
	}
	db.Model(&model.Answer{}).Create(&TestAnswer)
	resp := &proto.AcreateResp{}
	resp.OK = true
	return resp, nil
}
