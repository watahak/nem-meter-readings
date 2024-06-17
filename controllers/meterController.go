package meterController

import (
	"encoding/csv"
	"nem12/db"
	"nem12/pkg/parser"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadMeterReadings godoc
// @Summary Upload meter readings via CSV
// @Description Upload meter readings by uploading a CSV file
// @Tags meter_readings
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Router /meter-readings/upload [post]
func UploadMeterReadings(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to open file"})
		return
	}

	defer f.Close()

	reader := csv.NewReader(f)

	meterReadings, err := parser.ParseRecord(reader)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res := db.DB.Create(meterReadings)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, res.Error)
		return
	}

	c.JSON(http.StatusAccepted, meterReadings)
}
