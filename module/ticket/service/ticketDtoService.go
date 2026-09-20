package service

import (
	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
)

type TicketCreateDto struct {
	Title       string               `validate:"required,max=50"`
	Description string               `validate:"required,max=500"`
	Priority    model.TicketPriority `validate:"required,oneof=AVERAGE HIGH CRITICAL"`
}

type TicketChangeStatusDto struct {
	Status model.TicketStatus `validate:"required,oneof=RESOLVED CLOSED"`
}
