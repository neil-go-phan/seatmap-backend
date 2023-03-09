package handler

import (
	"net/http"
	"seatmap-backend/services"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	handler services.RoleService
}

func NewRoleHandler(handler services.RoleService) *RoleHandler {
	roleHandler := &RoleHandler{
		handler: handler,
	}
	return roleHandler
}

func (roleHandler *RoleHandler) GetRoles(c *gin.Context) {
	roles, err := roleHandler.handler.ListRole()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Bad request"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (roleHandler *RoleHandler) ValidateRole(c *gin.Context) {
	role := c.Query("role")
	err := roleHandler.handler.Validate(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Role invalid"})
		c.Abort()
		return
	}
	c.Next()
}