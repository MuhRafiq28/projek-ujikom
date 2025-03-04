package controllers

import (
    "net/http"
    "backend/database"
    "backend/models"
    "fmt"
    "strconv"
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


func UpdateSiswa(c *gin.Context) {
	// Mengambil id dari URL parameter
	id := c.Param("id")
	
	// Menambahkan log untuk memastikan ID yang diterima
	fmt.Println("ID yang diterima:", id)

	// Mengonversi id ke uint (sesuaikan dengan tipe data ID di model Siswa)
	idUint, err := strconv.ParseUint(id, 10, 32) // Parse ke uint (32-bit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var siswa models.Siswa

	// Mencari siswa berdasarkan ID
	if err := database.DB.First(&siswa, uint(idUint)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Siswa dengan ID " + id + " tidak ditemukan"})
		return
	}

	// Bind data JSON yang dikirimkan oleh client ke struct siswa
	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log data siswa yang diterima
	fmt.Println("Data siswa yang diterima:", siswa)

	// Update siswa ke database
	if err := database.DB.Save(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data siswa"})
		return
	}

	// Kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Data siswa berhasil diperbarui", "data": siswa})
}

func DeleteSiswa(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Siswa{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus siswa"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Siswa berhasil dihapus"})
}
