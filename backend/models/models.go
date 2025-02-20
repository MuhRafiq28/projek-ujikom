package models

import (
    "gorm.io/gorm"  // Pastikan Anda mengimpor GORM
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"not null" json:"password"` 
	Role     string `gorm:"default:user" json:"role"`
}

type Schedule struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Day   string `json:"day"`
	Guru  string `json:"Guru"`
	Time  string `json:"time"`
	Notes string `json:"notes"`
}

type Guru struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama      string `json:"nama"`
	KodeGuru  string `json:"kodeguru"`
	MapelGuru string `json:"mapelguru"`
}

type Izin struct {
	gorm.Model
	Nama       string    `json:"nama" binding:"required"`
	Status     string    `json:"status" binding:"required"`
	Alasan     string    `json:"alasan" binding:"required"`
	WaktuMasuk   *time.Time `json:"waktu_masuk"` // pointer to time.Time
    WaktuKeluar  *time.Time `json:"waktu_keluar"`
}


type Siswa struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    NIS       string    `json:"nis"`
    Nama      string    `json:"nama"`
    Jurusan   string    `json:"jurusan"`
    Kelas     string    `json:"kelas"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}


type Absensi struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	SiswaID   uint      `json:"siswa_id"`
	Siswa     Siswa `json:"siswa" gorm:"foreignKey:SiswaID"`
	Status    string    `json:"status"`
	Tanggal   string    `json:"tanggal"` // Pastikan formatnya "YYYY-MM-DD"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}







