package model

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	Answerer    string
	Message     string `form:"message" json:"message" binding:"required"`
	Question_id uint   `form:"question_id" json:"question_id" binding:"required"`
	Comments    []Comment
}

type Question struct {
	gorm.Model
	Questioner string
	Message    string `form:"message" json:"message" binding:"required"`
	Answers    []Answer
}

type Comment struct {
	gorm.Model
	Commenter string
	Message   string `form:"message" json:"message" binding:"required"`
	Answer_id uint   `form:"id" json:"id" binding:"required"`
}

type QuestionInfo struct {
	Questioner string
	Message    string
}

type QuestionModify struct {
	ID      uint   `form:"id" json:"id" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
}

type AnswerModify struct {
	ID      uint   `form:"id" json:"id" binding:"required"`
	Message string `form:"message" json:"message" binding:"required"`
}

type Delete struct {
	ID uint `form:"id" json:"id" binding:"required"`
}
