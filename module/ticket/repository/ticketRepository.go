package repository

import (
	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (t *TicketRepository) Save(ticket *model.TicketModel) error {
	err := t.db.Create(&ticket).Error

	if err != nil {
		return err
	}
	return nil
}

func (t *TicketRepository) FindAll() ([]*model.TicketModel, error) {
	var tickets []*model.TicketModel

	err := t.db.
		Preload("User").
		Preload("Attendant").
		Find(&tickets).Error

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (t *TicketRepository) FindByID(id string) (*model.TicketModel, error) {
	var ticket *model.TicketModel

	err := t.db.Where("id = ?", id).First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, nil

}

func (t *TicketRepository) Update(ticket *model.TicketModel) error {
	var ticketEntity *model.TicketModel

	err := t.db.Where("id = ?", ticket.ID).First(&ticketEntity).Error

	if err != nil {
		return err
	}

	t.db.Model(&ticketEntity).Updates(ticket)

	return nil
}

func (t *TicketRepository) Delete(id string) error {
	var ticket *model.TicketModel

	err := t.db.Where("id = ?", id).First(&ticket).Error
	if err != nil {
		return err
	}

	t.db.Delete(&ticket)

	return nil
}

func (t *TicketRepository) FindByTitle(title string) (*model.TicketModel, error) {
	var ticket *model.TicketModel

	err := t.db.Where("title LIKE ?", "%"+title+"%").First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (t *TicketRepository) ListMyTickets(userId string) ([]*model.TicketModel, error) {
	var tickets []*model.TicketModel

	err := t.db.
		Preload("User").
		Preload("Attendant").
		Where("ticket_models.user_id = ?", userId).
		Find(&tickets).Error

	if err != nil {
		return nil, err
	}

	return tickets, nil

}
