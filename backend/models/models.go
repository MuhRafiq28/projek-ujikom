package models

import (
    
	"time"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Jurusan  string `json:"jurusan"` // Menambahkan field jurusan
	Kelas    string `json:"kelas"`   // Menambahkan field kelas
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

// Struct untuk response


type Izin struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Nama        string     `json:"nama"`
	Status      string     `json:"status"`
	Alasan      string     `json:"alasan"`
	Jurusan     string     `json:"jurusan"`
	Kelas       string     `json:"kelas"`
	WaktuMasuk  *time.Time `json:"waktu_masuk"`
	WaktuKeluar *time.Time `json:"waktu_keluar"`
	
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


type Jadwal struct {
    ID              uint           `json:"id"`
    Jurusan         string         `json:"jurusan"`
    Kelas           string         `json:"kelas"`
    Hari            string         `json:"hari"`
    JamMulai        string         `json:"jam_mulai"`
    JamSelesai      string         `json:"jam_selesai"`
    MataPelajaran   MataPelajaran  `json:"mata_pelajaran" gorm:"foreignKey:MataPelajaranID;references:ID"`
    MataPelajaranID uint           `json:"mata_pelajaran_id"`
    Guru            string         `json:"guru"`
}

type MataPelajaran struct {
    ID   uint   `json:"id"`
    Nama string `json:"nama"`
}

