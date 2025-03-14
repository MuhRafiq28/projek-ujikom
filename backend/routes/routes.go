package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Routes untuk login dan register
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)

		// Routes untuk schedules
		api.GET("/schedules", controllers.GetAllSchedules)
		api.GET("/schedules/:id", controllers.GetScheduleByID)
		api.POST("/schedules", controllers.CreateSchedule)
		api.PUT("/schedules/:id", controllers.UpdateSchedule)
		api.DELETE("/schedules/:id", controllers.DeleteSchedule)

		// Routes untuk guru
		api.POST("/guru", controllers.CreateGuru)
		api.GET("/guru", controllers.GetGurus)
		api.GET("/guru/:id", controllers.GetGuruByID)
		api.PUT("/guru/:id", controllers.UpdateGuru)
		api.DELETE("/guru/:id", controllers.DeleteGuru)

		// Routes untuk izin
		api.POST("/izin", controllers.CreateIzin)        // Menambahkan izin
		api.GET("/izin", controllers.GetAllIzin)          // Mendapatkan semua izin
		api.DELETE("/izin/:id", controllers.DeleteIzin)   // Menghapus izin berdasarkan ID
		api.PUT("/izin/:id/konfirmasi", controllers.KonfirmasiMasuk) // Konfirmasi izin

		// Routes untuk siswa
		api.GET("/siswa", controllers.GetSiswa)           // Endpoint untuk get siswa dengan filter
		api.POST("/siswa", controllers.CreateSiswa)       // Endpoint untuk create siswa
		api.POST("/siswa/batch", controllers.CreateSiswaBatch) // Endpoint untuk create siswa batch
        api.PUT("/siswa/:id", controllers.UpdateSiswa)
		api.DELETE("/siswa/:id", controllers.DeleteSiswa)
		



		// Routes untuk absensi & RekapAbsensi
		api.POST("/absensi", controllers.CreateAbsensi)
		api.GET("/rekap-absensi/nama/:nama", controllers.GetRekapAbsensiByNama)
		api.DELETE("/rekap-absensi/:id", controllers.DeleteAbsensiByID)   
		
		

	

     
	}
}
