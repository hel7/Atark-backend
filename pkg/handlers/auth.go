package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) signUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

func (h *Handlers) signIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User signed in successfully"})
}
