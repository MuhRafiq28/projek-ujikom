package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"backend/models"
	"backend/database"
)

// Get all Gurus
func GetGurus(c *gin.Context) {
	var gurus []models.Guru

	if err := database.DB.Find(&gurus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gurus)
}

// Get Guru by ID (Perbaikan: Ambil dari database)
func GetGuruByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var guru models.Guru
	if err := database.DB.First(&guru, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guru not found"})
		return
	}

	c.JSON(http.StatusOK, guru)
}

// Create a new Guru (Perbaikan: Simpan ke database)
func CreateGuru(c *gin.Context) {
	var guru models.Guru
	if err := c.ShouldBindJSON(&guru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&guru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, guru)
}

// Update Guru by ID (Perbaikan: Update langsung di database)
func UpdateGuru(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var guru models.Guru
	if err := database.DB.First(&guru, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guru not found"})
		return
	}

	if err := c.ShouldBindJSON(&guru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&guru)
	c.JSON(http.StatusOK, guru)
}

// Delete Guru by ID (Perbaikan: Hapus dari database)
func DeleteGuru(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var guru models.Guru
	if err := database.DB.First(&guru, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guru not found"})
		return
	}

	database.DB.Delete(&guru)
	c.JSON(http.StatusOK, gin.H{"message": "Guru deleted successfully"})
}
