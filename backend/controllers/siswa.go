package controllers

import (
    "net/http"
    "backend/database"
    "backend/models"

    "github.com/gin-gonic/gin"
)

func GetSiswa(c *gin.Context) {
    var siswa []models.Siswa
    jurusan := c.Query("jurusan")
    kelas := c.Query("kelas")

    query := database.DB

    if jurusan != "" {
        query = query.Where("jurusan = ?", jurusan)
    }
    if kelas != "" {
        query = query.Where("kelas = ?", kelas)
    }

    query.Find(&siswa)

    c.JSON(http.StatusOK, siswa)
}


func CreateSiswa(c *gin.Context) {
	var siswa models.Siswa

	// Bind JSON ke struct Siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan ke database
	if err := database.DB.Create(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan siswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Siswa berhasil ditambahkan", "siswa": siswa})
}

func CreateSiswaBatch(c *gin.Context) {
    var siswaList []models.Siswa
    if err := c.ShouldBindJSON(&siswaList); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Create(&siswaList) // Masukkan semua siswa ke database
    c.JSON(http.StatusOK, gin.H{"message": "Data siswa berhasil ditambahkan", "data": siswaList})
}



