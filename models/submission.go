package models

import (
	"time"
)

type Submission struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SourceCode   string    `gorm:"type:text;not null" json:"sourceCode"`
	ResultStatus string    `gorm:"type:TEXT CHECK(result_status IN ('Running', 'Failed', 'Success'));not null" json:"resultStatus"`
	ProblemID    uint      `gorm:"not null" json:"problemId"`
	Problem      Problem   `gorm:"foreignKey:ProblemID" json:"-"`
	UserID       uint      `gorm:"not null" json:"userId"`
	User         User      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func GetSubmissionsByProblemID(problemID uint) ([]Submission, error) {
	var submissions []Submission
	result := DB.Where("problem_id = ?", problemID).Find(&submissions)
	return submissions, result.Error
}

func GetSubmissionsByProblemIDAndUserID(problemID uint, userID uint) ([]Submission, error) {
	var submissions []Submission
	result := DB.Where("problem_id = ? AND user_id = ?", problemID, userID).Find(&submissions)
	return submissions, result.Error
}

func CreateSubmission(submission *Submission) error {
	result := DB.Create(submission)
	return result.Error
}
