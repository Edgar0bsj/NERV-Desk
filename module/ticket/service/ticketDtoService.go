package service

import "github.com/edgar0bsj/nerv-desk/module/ticket/model"

type TicketCreateDto struct {
	Title       string               `validate:"required,max=50"`
	Description string               `validate:"required,max=500"`
	Priority    model.TicketPriority `validate:"required,oneof=AVERAGE HIGH CRITICAL"`
}

// type TicketUpdateDto struct {
// 	ID           string `validate:"required,uuid4"`
// 	Attendant_id string
// 	Title        string               `validate:"required,max=50"`
// 	Description  string               `validate:"required,max=500"`
// 	Status       model.TicketStatus   `validate:"required,oneof=OPEN IN_PROGRESS RESOLVED CLOSED"`
// 	Priority     model.TicketPriority `validate:"required,oneof=AVERAGE HIGH CRITICAL"`
// }
