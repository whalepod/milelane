package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/whalepod/milelane/app/domain"
	"github.com/whalepod/milelane/app/domain/repository"

	"github.com/whalepod/milelane/app/infrastructure"
)

// DeviceCreateJSON is struct for bind params to create device.
type DeviceCreateJSON struct {
	DeviceID   string `json:"device_id" binding:"required,min=1,max=36"`
	DeviceType string `json:"device_type" binding:"required"`
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

	device, err := d.Create(j.DeviceID, j.DeviceType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *device)
}
