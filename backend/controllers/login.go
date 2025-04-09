package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"
	"os" 
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("rahasia")) 

// Login user
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,  // Role yang diambil dari user
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token": tokenString,
		"username": user.Username,
		"id": user.ID,
		"role":  user.Role, // Menambahkan role pada response
	})
}


// Register user
func Register(c *gin.Context) {
    var input models.User

    // Menerima input dari body request
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi apakah username sudah ada
    var existingUser models.User
    if err := database.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    // Validasi role yang diizinkan
    if input.Role != "user" && input.Role != "siswa" && input.Role != "admin" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
        return
    }

    // Validasi jurusan dan kelas hanya untuk role siswa
    if input.Role == "siswa" && (input.Jurusan == "" || input.Kelas == "") {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Jurusan dan kelas harus diisi untuk siswa"})
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    input.Password = string(hashedPassword)

    // Menyimpan user baru
    if err := database.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "User registered successfully",
        "username": input.Username,
        "role": input.Role,
        "jurusan": input.Jurusan,  // Menampilkan jurusan pada response
        "kelas": input.Kelas,      // Menampilkan kelas pada response
    })
}

