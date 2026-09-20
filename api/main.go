package main

import (
	"log"

	ticketDatabase "github.com/edgar0bsj/nerv-desk/module/ticket/database"
	ticketHandler "github.com/edgar0bsj/nerv-desk/module/ticket/handler"
	ticketRepository "github.com/edgar0bsj/nerv-desk/module/ticket/repository"
	ticketService "github.com/edgar0bsj/nerv-desk/module/ticket/service"
	"github.com/edgar0bsj/nerv-desk/module/user/database"
	userhandler "github.com/edgar0bsj/nerv-desk/module/user/handler"
	"github.com/edgar0bsj/nerv-desk/module/user/repository"
	"github.com/edgar0bsj/nerv-desk/module/user/service"
	"github.com/edgar0bsj/nerv-desk/router"
)

func main() {
	// User Core
	userDb, _ := database.New()
	userRepo := repository.New(userDb)
	userSvc := service.New(userRepo)
	userHandler := userhandler.New(userSvc)

	// Ticket Core
	ticketDb, _ := ticketDatabase.New()
	ticketRepo := ticketRepository.New(ticketDb)
	ticketSvc := ticketService.New(ticketRepo)
	ticketHandler := ticketHandler.New(ticketSvc)

	r := router.SetupRouter(userHandler, ticketHandler)

	log.Println("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
