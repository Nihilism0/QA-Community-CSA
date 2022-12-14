package dao

import (
	"CSAwork/global"
	"log"
)

func SelectQuestion(id string) bool {
	flag, _ := global.RedisDb.SIsMember("questionids", id).Result()
	return flag
}

func Praiseadd(id string, username string) {
	global.RedisDb.SAdd(id, username)
}
func SelectPraiseuser(id, username string) bool {
	flag, _ := global.RedisDb.SIsMember(id, username).Result()
	return flag
}
func CancelPraise(id string, username string) {
	global.RedisDb.SRem(id, username)
}

func SeeQpraise(id string) int64 {
	result, err := global.RedisDb.SCard(id).Result()
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
