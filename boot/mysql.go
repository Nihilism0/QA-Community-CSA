package boot

import (
	"CSAwork/global"
	"CSAwork/model"
)

func MysqlSetUp() {
	global.GlobalDb1.AutoMigrate(&model.Comment{}, &model.Answer{}, &model.Question{})
}
