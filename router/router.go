package router

import (
	middle "github.com/edgar0bsj/nerv-desk/middleware"
	"github.com/edgar0bsj/nerv-desk/module/ticket/handler"
	userhandler "github.com/edgar0bsj/nerv-desk/module/user/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *userhandler.UserHandler, ticketHandler *handler.TicketHandler) *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", userHandler.UserLogin)
	}

	userRoutes := r.Group("/user", middle.RequireUserRole())
	{
		userRoutes.GET("/ticket", ticketHandler.FindAllTicketsUser)
		userRoutes.POST("/ticket", ticketHandler.CreateTicket)
		userRoutes.DELETE("/ticket/:id", ticketHandler.DeleteTicket)
	}

	return r
}
