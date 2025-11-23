package models

import "time"

type Answer struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	QuestionID uint      `json:"question_id"`
	UserID     string    `gorm:"type:varchar(36);not null" json:"user_id"`
	Text       string    `gorm:"type:text;not null" json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
