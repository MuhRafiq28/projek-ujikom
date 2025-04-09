package controllers

import (
	"backend/database"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetJadwal mengambil jadwal sesuai jurusan dan kelas user yang login
func GetJadwal(c *gin.Context) {
	// Ambil user dari context yang sudah diset di middleware
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan di context"})
		return
	}

	// Konversi ke model User
	user, ok := userData.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data user"})
		return
	}

	// Query awal untuk mengambil jadwal berdasarkan jurusan & kelas user
	query := database.DB.
		Joins("LEFT JOIN mata_pelajarans ON mata_pelajarans.id = jadwals.mata_pelajaran_id").
		Where("jadwals.jurusan = ? AND jadwals.kelas = ?", user.Jurusan, user.Kelas)

	// Tambahkan filter jika jurusan PH dan kelas 10A
	if user.Jurusan == "PH" && user.Kelas == "10A" {
		query = query.Where("jadwals.mata_pelajaran_id IN (?)", []uint{1, 3, 5})
	}

	// Debug: Cek query yang akan dieksekusi
	fmt.Println(query.Statement.SQL.String())

	// Ambil data jadwal
	var jadwal []models.Jadwal
	if err := query.Find(&jadwal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil jadwal"})
		return
	}

	// Kirim hasil dalam format JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "Jadwal ditemukan",
		"jadwal":  jadwal,
	})
}
