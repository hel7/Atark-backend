package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "UserID"
)

func (h *Handlers) userIndetity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("User id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("User id not valid type")
	}

	return idInt, nil
}
func (h *Handlers) adminRequired(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	user, err := h.services.GetUserByID(userID)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "User not found")
		return
	}

	if user.Role != "Admin" {
		newErrorResponse(c, http.StatusForbidden, "Admin role required")
		return
	}

	c.Next()
}
