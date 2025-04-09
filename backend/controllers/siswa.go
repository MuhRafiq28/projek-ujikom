package controllers

import (
    "net/http"
    "backend/database"
    "backend/models"
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// Mendapatkan daftar siswa berdasarkan filter jurusan dan kelas
// Mendapatkan daftar siswa berdasarkan filter jurusan, kelas, dan nama
func GetSiswa(c *gin.Context) {
    var siswa []models.Siswa
    jurusan := c.Query("jurusan")
    kelas := c.Query("kelas")
    nama := c.Query("nama") // ← ambil query nama

    query := database.DB
    if jurusan != "" {
        query = query.Where("jurusan = ?", jurusan)
    }
    if kelas != "" {
        query = query.Where("kelas = ?", kelas)
    }
    if nama != "" {
        query = query.Where("nama LIKE ?", "%"+nama+"%")
    }

    query.Find(&siswa)
    c.JSON(http.StatusOK, siswa)
}


// Menambahkan satu siswa dengan validasi
func CreateSiswa(c *gin.Context) {
    var siswa models.Siswa

    if err := c.ShouldBindJSON(&siswa); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if siswa.NIS == "" || siswa.Nama == "" || siswa.Jurusan == "" || siswa.Kelas == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "NIS, Nama, Jurusan, dan Kelas harus diisi"})
        return
    }

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

    // Cek jika ada data kosong dalam request
    nisSet := make(map[string]bool)
    var nisList []string

    for _, siswa := range siswaList {
        if siswa.NIS == "" || siswa.Nama == "" || siswa.Jurusan == "" || siswa.Kelas == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "NIS, Nama, Jurusan, dan Kelas harus diisi"})
            return
        }

        // Cek duplikasi dalam request JSON
        if nisSet[siswa.NIS] {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Terdapat NIS yang duplikat dalam input", "nis": siswa.NIS})
            return
        }
        nisSet[siswa.NIS] = true
        nisList = append(nisList, siswa.NIS)
    }

    // Periksa apakah ada NIS yang sudah terdaftar di database
    var existingSiswa []models.Siswa
    database.DB.Where("nis IN ?", nisList).Find(&existingSiswa)

    existingNisSet := make(map[string]bool)
    for _, siswa := range existingSiswa {
        existingNisSet[siswa.NIS] = true
    }

    // Filter siswa yang belum ada di database
    var newSiswaList []models.Siswa
    for _, siswa := range siswaList {
        if existingNisSet[siswa.NIS] {
            continue // Lewati siswa yang sudah ada
        }
        newSiswaList = append(newSiswaList, siswa)
    }

    // Jika semua NIS sudah ada, tidak perlu insert
    if len(newSiswaList) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Semua NIS sudah terdaftar", "existing": existingSiswa})
        return
    }

    // Simpan siswa yang baru
    if err := database.DB.Create(&newSiswaList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan data siswa"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Data siswa berhasil ditambahkan",
        "added":   newSiswaList,
        "skipped": existingSiswa,
    })
}
// Memperbarui data siswa dengan validasi
func UpdateSiswa(c *gin.Context) {
    id := c.Param("id")
    fmt.Println("ID yang diterima:", id)

    idUint, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
        return
    }

    var siswa models.Siswa
    if err := database.DB.First(&siswa, uint(idUint)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Siswa dengan ID " + id + " tidak ditemukan"})
        return
    }

    var input models.Siswa
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fmt.Println("Mengecek apakah NIS sudah digunakan:", input.NIS)

    var existingSiswa models.Siswa
    err = database.DB.Where("nis = ? AND id != ?", input.NIS, idUint).First(&existingSiswa).Error

    if err != nil && err != gorm.ErrRecordNotFound {
        fmt.Println("Error saat mengecek NIS:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat validasi NIS"})
        return
    }

    if err == nil {
        fmt.Println("NIS sudah digunakan oleh siswa lain:", existingSiswa.NIS)
        c.JSON(http.StatusBadRequest, gin.H{"error": "NIS sudah digunakan oleh siswa lain"})
        return
    }

    siswa.Nama = input.Nama
    siswa.NIS = input.NIS
    siswa.Kelas = input.Kelas
    siswa.Jurusan = input.Jurusan

    if err := database.DB.Save(&siswa).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data siswa"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data siswa berhasil diperbarui", "data": siswa})
}

// Menghapus data siswa berdasarkan ID
func DeleteSiswa(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Siswa{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus siswa"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Siswa berhasil dihapus"})
}

// Mendapatkan data pengguna berdasarkan ID
func GetUserData(c *gin.Context) {
    userID := c.Param("id")

    var user models.User
    if err := database.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "jurusan": user.Jurusan,
        "kelas":   user.Kelas,
    })
}
