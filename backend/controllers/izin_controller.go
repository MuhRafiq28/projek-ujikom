package controllers

import (
    "backend/database"
    "backend/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "time"
)

// CreateIzin - Menambahkan izin baru
func CreateIzin(c *gin.Context) {
    var izin models.Izin

    if err := c.ShouldBindJSON(&izin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Panggil fungsi dari models
	izin.SetWaktuByStatus()

    // Mengambil waktu sekarang dalam zona waktu Asia/Jakarta
jakarta, _ := time.LoadLocation("Asia/Jakarta")
now := time.Now().In(jakarta)

    

switch izin.Status {
case "Keluar", "Pulang Lebih Cepat":
    izin.WaktuKeluar = &now
    izin.WaktuMasuk = nil
case "Masuk", "Terlambat":
    izin.WaktuMasuk = &now
    izin.WaktuKeluar = nil
default:
    izin.WaktuKeluar = nil
    izin.WaktuMasuk = nil
}


    // Simpan data ke database
    if err := database.DB.Create(&izin).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan izin"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Izin berhasil ditambahkan", "data": izin})
}

// GetAllIzin - Mengambil semua data izin
func GetAllIzin(c *gin.Context) {
    var izins []models.Izin
    if err := database.DB.Find(&izins).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data izin"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": izins})
}

// DeleteIzin - Menghapus izin berdasarkan ID
func DeleteIzin(c *gin.Context) {
    id := c.Param("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
        return
    }

    var izin models.Izin
    if err := database.DB.First(&izin, idInt).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Izin tidak ditemukan"})
        return
    }

    if err := database.DB.Delete(&izin).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus izin"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Izin berhasil dihapus"})
}


func KonfirmasiMasuk(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var izin models.Izin
	if err := database.DB.First(&izin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data izin tidak ditemukan"})
		return
	}

	if izin.Status != "Keluar" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Izin belum keluar atau sudah dikonfirmasi"})
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
now := time.Now().In(loc)

izin.Status = "Sudah Kembali"
izin.WaktuMasuk = &now

	if err := database.DB.Save(&izin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan perubahan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Siswa berhasil dikonfirmasi masuk",
		"data":    izin,
	})
}
