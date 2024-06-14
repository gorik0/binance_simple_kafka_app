package handler

import (
	"bina/internal/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CreateAccRequest struct {
	Currency string `json:"currency" binding:"required"`
}
type CreateAccResponse struct {
	Id       int    `json:"id" `
	Balance  int64  `json:"balance" `
	Currency string `json:"currency" `
	CreateAt string `json:"create_at"`
}
func (h *Handler) createAcc(c *gin.Context){

	var req CreateAccRequest
	err := c.BindJSON(&req)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,fmt.Sprintf("Bad request :::%s",err))


		return
	}

	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error ::: %s",err))
		return
	}

	acc := core.Account{
		Id:       0,
		UserId:   id,
		Balance:  0,
		Currency: req.Currency,
	}
	accId, err := h.service.Account.CreateAccount(&acc)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Counlt create acc ::: %s ",err))
		return
	}

	accToReturn := CreateAccResponse{
		Id:       accId,
		Balance:  acc.Balance,
		Currency: acc.Currency,
		CreateAt: time.Now().String(),
	}
	c.JSON(http.StatusOK,accToReturn)

}

type GetAccountRequest struct {
	Id int `uri:"id" binding:"required,min=1"`
}
func (h *Handler) getAcc(c *gin.Context){

	var req GetAccountRequest
	err := c.BindUri(req)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"Bad request")
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get user id)  ::: %s",err))
		return
	}

	account, err := h.service.Account.GetAccountById(req.Id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get account by  id from repo )  ::: %s",err))
		return
	}
	if account.UserId != userID{
		newErrorResponse(c,http.StatusForbidden,"Try to acces foreign account !!!")
		return

	}
	c.JSON(http.StatusOK,account)

	
}
func (h *Handler) getAccountsById(c *gin.Context){

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get user id)  ::: %s",err))
		return
	}

	accounts, err := h.service.Account.GetAccounts(userID)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get accounts from repo )  ::: %s",err))
		return
	}
	c.JSON(http.StatusOK,accounts)
}

type UpdateAccRequest struct {
	Balance int64 `json:"balance" binding:required`
	Currency string `json:"currency" binding:required`
}
func (h *Handler) updateAcc(c *gin.Context){



	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"Invalid param(id) type")
		return
	}


	var req UpdateAccRequest
	err = c.BindJSON(req)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,fmt.Sprintf("Bad request(body) :::  %s",err))
		return
	}

	userID, err := getUserId(c)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get user id)  ::: %s",err))
		return
	}

	account, err := h.service.Account.GetAccountById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error (to get account by  id from repo )  ::: %s",err))
		return
	}
	if account.UserId != userID{
		newErrorResponse(c,http.StatusForbidden,"Try to acces foreign account !!!")
		return

	}

	if account.Currency!=req.Currency {
		newErrorResponse(c,http.StatusForbidden,"Cutrrency doesn't match !!!")
		return
	}
	if account.Balance<0 {
		newErrorResponse(c,http.StatusBadRequest,"Balance can't be negative !!!")
		return
	}


	account.Balance = req.Balance

	err = h.service.Account.UpdateAccount(account)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error while update!!!  ::: %s",err))

		return
	}
	c.JSON(http.StatusOK,account)
}