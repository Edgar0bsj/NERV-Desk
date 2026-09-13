package service

import (
	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
)

type TicketInterfaceService interface {
	SaveTicket(ticketDto *TicketCreateDto) error
	FindAllTickets() ([]*model.TicketModel, error)
	FindByIdTicket(id string) (*model.TicketModel, error)
	UpdateTicket(ticketDto *TicketUpdateDto) error
	DeleteTicket(id string) error
	FindByTitle(title string) (*model.TicketModel, error)
	FindAllByUserId(user_id string) ([]*model.TicketModel, error)
}
