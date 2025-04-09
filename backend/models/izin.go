package models

import "time"

func (izin *Izin) SetWaktuByStatus() {
	now := time.Now()
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
}
