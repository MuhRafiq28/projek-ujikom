package controllers

import (
	"net/http"
	"backend/database"
	"backend/models"
	"strconv" // pastikan package strconv diimpor
	"github.com/gin-gonic/gin"
)

// ✅ Fungsi untuk Menambahkan Absensi dalam jumlah banyak
func CreateAbsensi(c *gin.Context) {
	var absensiList []models.Absensi

	// Coba parsing JSON dari frontend
	if err := c.ShouldBindJSON(&absensiList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek validitas siswa_id dan simpan ke database
	for _, absensi := range absensiList {
		var siswa models.Siswa
		if err := database.DB.First(&siswa, absensi.SiswaID).Error; err != nil {
			// Menggunakan strconv.Itoa untuk konversi uint ke string
			c.JSON(http.StatusBadRequest, gin.H{"error": "Siswa dengan ID " + strconv.Itoa(int(absensi.SiswaID)) + " tidak ditemukan"})
			return
		}

		// Simpan absensi ke database
		if err := database.DB.Create(&absensi).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan absensi"})
			return
		}
	}

	// Kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Absensi berhasil ditambahkan"})
}


// Fungsi untuk mengambil rekap absensi berdasarkan nama siswa
func GetRekapAbsensiByNama(c *gin.Context) {
    // Mengambil parameter nama dari URL
    nama := c.Param("nama")

    // Logika untuk mengambil data absensi berdasarkan nama siswa
    var absensi []models.Absensi
    err := database.DB.Joins("JOIN siswas ON siswas.id = absensis.siswa_id").
        Where("LOWER(siswas.nama) LIKE LOWER(?)", "%"+nama+"%").
        Preload("Siswa").
        Find(&absensi).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data absensi"})
        return
    }

    if len(absensi) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ada absensi untuk siswa dengan nama tersebut"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "success", "data": absensi})
}


// ✅ Fungsi untuk Mengambil Daftar Absensi (dengan filter jurusan & kelas)
func GetAbsensi(c *gin.Context) {
	var absensi []models.Absensi
	jurusan := c.Query("jurusan")
	kelas := c.Query("kelas")

	// Query awal dengan preload siswa
	query := database.DB.Preload("Siswa")

	// Gunakan Join agar bisa filter berdasarkan tabel siswa
	if jurusan != "" || kelas != "" {
		query = query.Joins("JOIN siswas ON siswas.id = absensis.siswa_id")
	}

	// Filter berdasarkan jurusan
	if jurusan != "" {
		query = query.Where("siswas.jurusan = ?", jurusan)
	}

	// Filter berdasarkan kelas
	if kelas != "" {
		query = query.Where("siswas.kelas = ?", kelas)
	}

	// Jalankan query
	if err := query.Find(&absensi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data absensi"})
		return
	}

	// Kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"data": absensi})
}


// Hapus absensi berdasarkan ID
func DeleteAbsensiByID(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&models.Absensi{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Data absensi tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Data absensi berhasil dihapus"})
}

func DeleteAbsensiByNama(c *gin.Context) {
    nama := c.Param("nama")

    // Cari siswa berdasarkan nama
    var siswa models.Siswa
    if err := database.DB.Where("nama = ?", nama).First(&siswa).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Siswa tidak ditemukan"})
        return
    }

    // Hapus absensi berdasarkan SiswaID
    result := database.DB.Where("siswa_id = ?", siswa.ID).Delete(&models.Absensi{})
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Tidak ada data absensi untuk siswa ini"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Data absensi siswa berhasil dihapus"})
}

// Hapus semua data absensi
func DeleteAllAbsensi(c *gin.Context) {
	result := database.DB.Exec("DELETE FROM absensis")

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Gagal menghapus semua data absensi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Semua data absensi berhasil dihapus"})
}


