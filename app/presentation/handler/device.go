package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/whalepod/milelane/app/domain"
	"github.com/whalepod/milelane/app/domain/repository"

	"github.com/whalepod/milelane/app/infrastructure"
)

// DeviceCreateJSON is struct for bind params to create device.
type DeviceCreateJSON struct {
	DeviceToken string `json:"device_token" binding:"required,min=1,max=36"`
	DeviceType  string `json:"device_type" binding:"required"`
}

// DeviceCreate is handler to create device.
func DeviceCreate(c *gin.Context) {
	deviceAccessor := repository.NewDevice(infrastructure.DB)
	d, _ := domain.NewDevice(deviceAccessor)

	var j DeviceCreateJSON
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	device, err := d.Create(j.DeviceToken, j.DeviceType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *device)
}
