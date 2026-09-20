package model

import (
	"errors"
	"time"

	"github.com/edgar0bsj/nerv-desk/module/user/model"
)

type TicketStatus string

const (
	StatusOpen      TicketStatus = "OPEN"
	StatusInProcess TicketStatus = "IN_PROGRESS"
	StatusResolved  TicketStatus = "RESOLVED"
	StatusClosed    TicketStatus = "CLOSED"
)

type TicketPriority string

const (
	PriorityAverage  TicketPriority = "AVERAGE"
	PriorityHigh     TicketPriority = "HIGH"
	PriorityCritical TicketPriority = "CRITICAL"
)

type TicketModel struct {
	ID string `gorm:"primaryKey;type:varchar(50)"`

	UserID string      `gorm:"not null"`
	User   *model.User `gorm:"foreignKey:UserID"`

	AttendantID string      `gorm:"default:null"`
	Attendant   *model.User `gorm:"foreignKey:AttendantID"`

	Title       string `gorm:"not null"`
	Description string
	Status      TicketStatus
	Priority    TicketPriority
	Created_at  time.Time
	Updated_at  time.Time
}

var (
	ErrTicketNotFound   = errors.New("Ticket Not Found")
	ErrInvalidTicket    = errors.New("Invalid ticket data")
	ErrSaveFailedTicket = errors.New("Error persisting ticket data")
)
