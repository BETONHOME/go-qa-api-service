package models

import "time"

type Question struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Text      string    `gorm:"type:text" json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Answers   []Answer  `json:"answers,omitempty" gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE;"`
}
