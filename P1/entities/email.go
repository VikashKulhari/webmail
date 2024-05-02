package entities

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	From         string    `gorm:"not null" json:"From"`
	To           string    `gorm:"not null" json:"To"`
	Subject      string    `gorm:"not null" json:"Subject"`
	Body         string    `gorm:"not null" json:"Body"`
	IsSuccess    bool      `gorm:"not null" json:"IsSuccess"`
	IsDraft      bool      `gorm:"not null" json:"IsDraft"`
	IsSpam       bool      `gorm:"not null" json:"IsSpam"`
	IsImportant  bool      `gorm:"not null" json:"IsImportant"`
	IsDeleted    bool      `gorm:"not null" json:"IsDeleted"`
	RepliesCount int       `gorm:"not null"`
	// Thread       []Replies `gorm:"foriegnKey:MailID;onDelete:CASCADE"`
}
type Replies struct {
	gorm.Model
	MailID   string
	ReplySeq int
	From     string
	To       string
	Body     string `grom:"not null"`
}

type EmailReq struct {
	From         string    `json:"from" validate:"required"`
	To           string    `json:"to"`
	Subject      string    `json:"subject" validate:"required"`
	Body         string    `json:"body" validate:"required"`
	IsDraft      bool      `json:"isDraft" validate:"required"`
	IsSuccess    bool      `json:"isSuccess,omitempty"`
	IsImportant  bool      `json:"isImportant,omitempty"`
	IsDeleted    bool      `json:"isDeleted,omitempty"`
	IsSpam       bool      `json:"isSpam,omitempty"`
	RepliesCount int       `json:"repliesCount,omitempty"`
	Thread       []Replies `json:"thread,omitempty"`
}
