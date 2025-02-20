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

    // Set waktu masuk ke sekarang jika statusnya "Masuk"
    if izin.Status == "Masuk" {
        now := time.Now()
        izin.WaktuMasuk = &now  // Menggunakan pointer ke time.Now()
    }

    // Set waktu keluar jika statusnya "Keluar"
    if izin.Status == "Keluar" {
        now := time.Now()
        izin.WaktuKeluar = &now // Menggunakan pointer ke time.Now()
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

    c.JSON(http.StatusOK, gin.H{"message": "Berhasil mengambil data izin", "data": izins})
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

// KonfirmasiMasuk - Mengubah status izin menjadi "Sudah Kembali"
func KonfirmasiMasuk(c *gin.Context) {
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

    // Ubah status menjadi "Sudah Kembali" dan set waktu masuk
    izin.Status = "Sudah Kembali"
    now := time.Now()
    izin.WaktuMasuk = &now // Menggunakan pointer ke time.Now()

    if err := database.DB.Save(&izin).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate izin"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Izin dikonfirmasi masuk", "data": izin})
}
