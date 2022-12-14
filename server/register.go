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

func Register() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50057")
	log.Println("已监听50057,作用是注册微服务")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //获取新服务示例
	proto.RegisterRegisterServer(s, &Registerserver{})
	// 开始处理
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Registerserver struct {
	proto.UnimplementedRegisterServer // 用于实现proto包里RegisterServer接口
}

func (s *Registerserver) mustEmbedUnimplementedRegisterServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Registerserver) Register(ctx context.Context, req *proto.UserReq) (*proto.UserResp, error) {
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
	db.Model(&model.User{}).Create(&model.User{
		Username: req.UserName,
		Password: req.PassWord,
	})
	resp := &proto.UserResp{}
	resp.OK = true
	return resp, nil
}
