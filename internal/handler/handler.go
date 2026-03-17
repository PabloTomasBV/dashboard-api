package handler

import (
	"dashboard-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service *service.DashboardService
}

func NewDashboardHandler(s *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: s}
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error ": "invalid id",
		})
		return
	}

	if id > 8 || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error ": "id out of range (1-8)",
		})
		return
	}

	result, err := h.service.GetDashboard(strconv.Itoa(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error ": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
