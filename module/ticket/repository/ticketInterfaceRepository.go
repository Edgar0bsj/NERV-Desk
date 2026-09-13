package repository

import "github.com/edgar0bsj/nerv-desk/module/ticket/model"

type TicketInterfaceRepository interface {
	Save(ticket *model.TicketModel) error
	FindAll() ([]*model.TicketModel, error)
	FindByID(id string) (*model.TicketModel, error)
	Update(ticket *model.TicketModel) error
	Delete(id string) error
	FindByTitle(title string) (*model.TicketModel, error)
	FindAllByUserId(user_id string) ([]*model.TicketModel, error)
}
