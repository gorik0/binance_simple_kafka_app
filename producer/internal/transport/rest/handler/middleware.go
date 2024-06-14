package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const bearerToken = "Bearer"

const authorizationHeader = "Authorization"

func (h *Handler) userIdentity(c *gin.Context){
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != bearerToken {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(UserIdKey, userId)
}

const UserIdKey = "userID"

func getUserId(c *gin.Context) (int, error) {

	value, ok := c.Get(UserIdKey)
	if  !ok {
		return -1,errors.New("No user!")
	}
	id,ok := value.(int)
	if !ok {
		return -1,errors.New("Id no valid type of!")

	}
	return id,nil

}
