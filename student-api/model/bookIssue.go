package model

import uuid "github.com/satori/go.uuid"

type BookIssue struct {
	Base
	BookID    uuid.UUID `gorm:"type:varchar(36)" json:"bookId"`
	StudentID uuid.UUID `gorm:"type:varchar(36)" json:"studentId"`
	Book      Book      `json:"book"`
	IssueDate string    `gorm:"type:datetime" json:"issueDate"`
	Returned  bool      `gorm:"type:tinyint" json:"returned"`
	Penalty   float64   `gorm:"type:decimal(10,2)" json:"penalty"`
}
