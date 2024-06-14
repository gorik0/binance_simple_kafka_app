package handler

import (
	"bina/internal/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpReponse struct {


	Id int `json:"id"`

}

type SingInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInReponse struct {

	Token string `json:"id"`

}
func (h *Handler) signIn(c *gin.Context) {

	var request SingInRequest

	err := c.BindJSON(&request)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"Invalid request!!!")
		return
	}

	token, err := h.service.Authorization.GenerateToken(request.Username, request.Password)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,fmt.Sprintf("Internal error !!! ::: %s",err))
		return
	}


	c.JSON(http.StatusOK,SignInReponse{Token: token})

}

func (h *Handler) signUp(c *gin.Context) {


	var request core.User

	err := c.BindJSON(&request)
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"Invalid request!!!")
		return
	}

	id, err := h.service.CreateUser(&request)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,"Internal error  !!!")

		return
	}
	c.JSON(http.StatusOK,SignUpReponse{Id: id})

}

