package handler

import (
	"bina/internal/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTransferRequest struct {
	FromAccountID int    `json:"from_account_id" binding:"required"`
	ToAccountID   int    `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
}
type CreateTransferResponse struct {

	Id int `json:"id"`
}
func (h *Handler) createTransfer(c *gin.Context){
	var req CreateTransferRequest
	err := c.BindJSON(req)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"Bad request")
		return
	}

	trans := core.Transfer{
		Id:            0,
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
	}
	id, err := h.service.Transfer.CreateTransfer(&trans)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Couldn't create transfer :::: %s",err))
		return
	}
	c.JSON(http.StatusOK,CreateTransferResponse{Id: id})

}

type GetTransferByIdRequest struct {
	Id int `uri:"id" binding:"required,min=1"`

}
func (h *Handler) getTransfer(c *gin.Context){

	var req GetTransferByIdRequest
	err := c.ShouldBindUri(req)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"Bad request")
		return
	}

	transfer, err := h.service.Transfer.GetTransferById(req.Id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Couldn't gte trans bu ID :::: %s",err))
		return
	}

	c.JSON(http.StatusOK,transfer)

}
func (h *Handler) getTransfersById(c *gin.Context){


	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get user id)  ::: %s",err))
		return
	}

	transfer, err := h.service.Transfer.GetTransfers(userID)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get account by  id from repo )  ::: %s",err))
		return
	}
	c.JSON(http.StatusOK,transfer)


}
