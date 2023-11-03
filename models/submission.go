package models

import "gorm.io/gorm"

type SubmissionStatus string

const (
	PendingReview SubmissionStatus = "PENDING"
	UnderReview   SubmissionStatus = "UNDER_REVIEW"
	Accepted      SubmissionStatus = "ACCEPTED"
	Rejected      SubmissionStatus = "REJECTED"
)

type Submission struct {
	gorm.Model
	Proposal string           `gorm:"column:proposal;"`
	Status   SubmissionStatus `gorm:"column:status;"`
}
