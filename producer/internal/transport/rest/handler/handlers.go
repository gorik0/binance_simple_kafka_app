package handler

import (

	"bina/internal/service"
	"github.com/gin-gonic/gin"
)


type Handler struct {
	service *service.Service

}



func NewHandler(service2 *service.Service)*Handler{
	return &Handler{service: service2}
}
func (h *Handler) InitRoutes() *gin.Engine  {

router:= gin.New()

auth:= router.Group("/auth")
{
	auth.POST("/sign-in",h.signIn)
	auth.POST("/sign-up",h.signUp)
}

api:= router.Group("/api",h.userIdentity)
{
	acc:= api.Group("/account")
	{
		acc.POST("/",h.createAcc)
		acc.GET("/",h.getAcc)
		acc.GET("/:id",h.getAccountsById)
		acc.PUT("/:id",h.updateAcc)
	}
	tra:= api.Group("/transfer")

	{
		tra.POST("/",h.createTransfer)
		tra.GET("/",h.getTransfer)
		tra.GET("/:id",h.getTransfersById)
	}

	coin:= api.Group("/coin")
	{
		coin.GET("/price",h.getPrice)
	}
}

	return router
}



