package controllers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"ms-live-sells/models"
	"ms-live-sells/services"
	"net/http"
)

type InstagramController struct {
	InstagramService *services.InstagramService
}

// StartMonitoring starts the live monitoring for a user
func (ctrl *InstagramController) StartMonitoring(c echo.Context) error {
	// Parse the request
	req := new(models.InstagramMonitorRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Validate the request
	if req.UserID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "userID is required"})
	}

	// Start monitoring
	go func() {
		err := ctrl.InstagramService.StartInstagramMonitoring(req.UserID, req.LiveID)
		if err != nil {
			log.Printf("Error starting monitoring for user %s: %v", req.UserID, err)
		}
	}()

	return c.JSON(http.StatusAccepted, map[string]string{"message": "monitoring started successfully"})
}
