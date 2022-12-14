package boot

import (
	g "CSAwork/global"
	"CSAwork/model"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func MysqlDBSetup() {
	config := g.Config.DataBase.Mysql
	db, err := gorm.Open(mysql.Open(config.GetDsn()))
	if err != nil {
		g.Logger.Fatal("initialize mysql failed.", zap.Error(err))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(g.Config.DataBase.Mysql.GetConnMaxIdleTime())
	sqlDB.SetConnMaxLifetime(g.Config.DataBase.Mysql.GetConnMaxLifeTime())
	sqlDB.SetMaxIdleConns(g.Config.DataBase.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(g.Config.DataBase.Mysql.MaxOpenConns)
	err = sqlDB.Ping()
	if err != nil {
		g.Logger.Fatal("connect to mysql db failed.", zap.Error(err))
	}
	g.GlobalDb1 = db
	g.GlobalDb1.AutoMigrate(&model.User{}, &model.Question{}, &model.Answer{})
	g.Logger.Info("initialize mysql successfully!")
}

func RedisSetup() {
	config := g.Config.DataBase.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	g.Ctx = ctx
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		g.Logger.Fatal("connect to redis instance failed.", zap.Error(err))
	}
	g.RedisDb = rdb
	g.Logger.Info("initialize redis client successfully!")
}
