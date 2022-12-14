package dao

import (
	"CSAwork/global"
	"CSAwork/model"
)

func FindQuestionSubmited(username string) []model.Question {
	var questions []model.Question
	global.GlobalDb1.Model(&model.Question{}).Preload("Answers").Preload("Comments").Where("Questioner = ?", username).Find(&questions)
	return questions
}

func FindAnswerSubmited(username string) []model.Answer {
	var answers []model.Answer
	global.GlobalDb1.Model(&model.Answer{}).Preload("Comments").Where("Answerer = ?", username).Find(&answers)
	return answers
}

func JudgeQuestion(username string, id int) bool {
	OK := false
	var question model.Question
	global.GlobalDb1.Model(&model.Question{}).Where("questioner = ? AND id = ?", username, id).First(&question)
	if question.Questioner == username {
		OK = true
	}
	return OK
}

func QuestionModify(message string, id int) {
	global.GlobalDb1.Model(&model.Question{}).Where("id = ?", id).Update("message", message)
}

func JudgeAnswer(username string, id int) bool {
	OK := false
	var answer model.Answer
	global.GlobalDb1.Model(&model.Answer{}).Where("answerer = ? AND id = ?", username, id).First(&answer)
	if answer.Answerer == username {
		OK = true
	}
	return OK
}

func AnswerModify(message string, id int) {
	global.GlobalDb1.Model(&model.Answer{}).Where("id = ?", id).Update("message", message)
}

func AnswerDelete(id int) {
	var answer model.Answer
	var comments []model.Comment
	global.GlobalDb1.Model(&model.Comment{}).Unscoped().Where("answer_id = ?", id).Delete(&comments)
	global.GlobalDb1.Model(&model.Answer{}).Unscoped().Where("id = ?", id).Delete(&answer)
}
func QuestionDelete(id int) {
	var question model.Question
	var answer []model.Answer
	var comments []model.Comment
	global.GlobalDb1.Model(&model.Comment{}).Unscoped().Where("answer_id = ?", id).Delete(&comments)
	global.GlobalDb1.Model(&model.Answer{}).Unscoped().Where("question_id = ?", id).Delete(&answer)
	global.GlobalDb1.Model(&model.Question{}).Unscoped().Where("id = ?", id).Delete(&question)
}
