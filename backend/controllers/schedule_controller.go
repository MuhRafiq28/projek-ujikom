package controllers

import (
	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/models"
	"net/http"
	"log"
)

// GetAllSchedules - Mengambil semua jadwal
func GetAllSchedules(c *gin.Context) {
	var schedules []models.Schedule

	// Query semua data dari database
	if err := database.DB.Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		log.Println("Error fetching schedules:", err)
		return
	}

	// Kirim respons sebagai JSON
	c.JSON(http.StatusOK, schedules)
}

// GetScheduleByID - Mengambil jadwal berdasarkan ID
func GetScheduleByID(c *gin.Context) {
	// Ambil parameter ID dari URL
	id := c.Param("id")
	var schedule models.Schedule

	// Cari jadwal berdasarkan ID
	if err := database.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		log.Println("Error finding schedule:", err)
		return
	}

	// Kirim respons sebagai JSON
	c.JSON(http.StatusOK, schedule)
}

// CreateSchedule - Menambahkan jadwal baru
func CreateSchedule(c *gin.Context) {
	var schedule models.Schedule

	// Decode body JSON ke struct
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan ke database
	if err := database.DB.Create(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	// Kirim respons
	c.JSON(http.StatusCreated, schedule)
}

// UpdateSchedule - Memperbarui jadwal berdasarkan ID
func UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	var schedule models.Schedule

	// Cari jadwal berdasarkan ID
	if err := database.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		log.Println("Error finding schedule:", err)
		return
	}

	// Decode body JSON ke struct
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pastikan field penting tidak kosong
	if schedule.Day == "" || schedule.Guru == "" || schedule.Time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Day, Teachers, and Time must be provided"})
		return
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		log.Println("Error updating schedule:", err)
		return
	}

	// Kirim respons
	c.JSON(http.StatusOK, schedule)
}

// DeleteSchedule - Menghapus jadwal berdasarkan ID
func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	var schedule models.Schedule

	// Cari jadwal berdasarkan ID
	if err := database.DB.First(&schedule, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		log.Println("Error finding schedule:", err)
		return
	}

	// Hapus jadwal dari database
	if err := database.DB.Delete(&schedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		log.Println("Error deleting schedule:", err)
		return
	}

	// Kirim respons
	c.Status(http.StatusNoContent)
}

