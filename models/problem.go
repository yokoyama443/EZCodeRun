package models

import (
	"time"
)

type Problem struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Title             string    `gorm:"not null" json:"title"`
	TimeLimit         int       `gorm:"not null" json:"timeLimit"`
	MemoryLimit       int       `gorm:"not null" json:"memoryLimit"`
	FreeFeedBackLimit int       `gorm:"not null" json:"freeFeedBackLimit"`
	Body              string    `gorm:"type:text;not null" json:"body"`
	TestCaseInput     string    `gorm:"type:text;not null" json:"testCaseInput"`
	TestCaseOutput    string    `gorm:"type:text;not null" json:"testCaseOutput"`
	Priority          int       `gorm:"not null" json:"priority"`
	CreatedAt         time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func GetAllProblems() ([]Problem, error) {
	var problems []Problem
	result := DB.Find(&problems)
	return problems, result.Error
}

func GetProblemByID(id uint) (Problem, error) {
	var problem Problem
	result := DB.First(&problem, id)
	return problem, result.Error
}
