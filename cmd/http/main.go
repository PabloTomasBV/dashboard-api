package main

import (
	"dashboard-api/internal/handler"
	"dashboard-api/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}

	service := service.NewDashboardService(httpClient, "https://dummyjson.com")
	handler := handler.NewDashboardHandler(service)

	router.GET("/dashboard/:id", handler.GetDashboard)
	router.Run(":8080")
}
