<template>
  <div class="container-fluid p-2">
    <div class="hero d-flex justify-content-between align-items-center">
      <div class="left-text">Academic School Internasional</div>
      <div class="right-img">
        <!-- Notifikasi muncul selama 2 detik di atas gambar -->
        <div v-if="showNotification" class="notification">
          Klik foto untuk membuka profil
        </div>
        <!-- Membungkus gambar dengan nuxt-link -->
        <nuxt-link to="/profile">
          <img src="/images/inspirasi-logo.png" alt="SMKN 1 Cisarua" />
        </nuxt-link>
      </div>
    </div>
    <div class="title">
      <h1>Kalian adalah generasi hebat yang sedang mempersiapkan masa depan. Jangan pernah meremehkan potensi diri sendiri! SMK bukan hanya tentang sekolah, tapi tentang bagaimana kalian bisa menjadi ahli di bidang yang kalian pilih.</h1>
      <h3>tetap semangat jadi ketua kelas</h3>
      <h4>{{ username }}</h4>
    </div>
    <div class="copyright">
      <p>© 2025 Muhammad Rafiq. All rights reserved.</p>
    </div>
    <div class="navbar">
      <nuxt-link to="/siswa/absen">
        <i class="fas fa-calendar-check"></i> Absen
      </nuxt-link>
      <span class="home">
        <nuxt-link to="">
          <i class="fas fa-home"></i> Home
        </nuxt-link>
      </span>
      <nuxt-link to="">
        <i class="fas fa-sign-out-alt"></i> Izin
      </nuxt-link>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showNotification: true, // Variabel untuk menampilkan notifikasi
      username: "", // Variabel untuk menyimpan username
    };
  },
  mounted() {
     // Mengambil username dari localStorage dan menyimpannya di variabel username
     const storedUsername = localStorage.getItem("username");
    if (storedUsername) {
      this.username = storedUsername;
    }

    // Menghilangkan notifikasi setelah 2 detik
    setTimeout(() => {
      this.showNotification = false;
    }, 2000); // 2000 ms = 2 detik
  },
};
</script>

<style>
.container-fluid {
  margin: none;
}
.hero {
  width: 100%;
}

.left-text {
  text-align: left;
  margin-left: 9px;
}

.right-img {
  text-align: right;
  position: relative; /* Agar notifikasi dapat ditempatkan di atas gambar */
}

img {
  width: 80px;
  height: auto;
}

.notification {
  position: absolute;
  top: -40px;
  left: 50%;
  transform: translateX(-50%);
  background: linear-gradient(45deg, #007bff, #00d2ff);
  color: white;
  padding: 12px 20px;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  opacity: 0;
  animation: fadeIn 1s ease-in-out forwards, fadeOut 1s 1.5s ease-out forwards;
}

/* Animasi untuk Fade In */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Animasi untuk Fade Out */
@keyframes fadeOut {
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
    transform: translateY(-20px);
  }
}

/* Title Section */
.title {
  text-align: center;
  margin-top: 100px;
  padding: 0 20px;
  animation: fadeIn 1.2s ease-in-out;
}

/* Desain untuk H1 */
.title h1 {
  font-size: 20px;
  line-height: 1.6;
  font-weight: 500;
  margin: 0;
  color: #333;
  text-transform: capitalize;
  letter-spacing: 1px;
  text-shadow: 1px 1px 10px rgba(0, 0, 0, 0.15);
  animation: slideUp 1.5s ease-out;
}

/* Desain untuk H3 */
.title h3 {
  font-size: 26px;
  margin-top: 20px;
  font-weight: bold;
  color: #007bff;
  background: linear-gradient(45deg, #007bff, #00d2ff);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  animation: fadeInText 2s ease-in-out;
}

/* Desain untuk H4 */
.title h4 {
  font-size: 18px;
  margin-top: 20px;
  font-weight: 600;
  color: #555;
}

/* Animasi Fade In */
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

/* Animasi Slide Up */
@keyframes slideUp {
  from {
    transform: translateY(50px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* Animasi Fade In untuk Teks */
@keyframes fadeInText {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Responsif untuk Mobile */
@media (max-width: 768px) {
  .title {
    margin-top: 70px;
  }

  .title h1 {
    font-size: 18px;
  }

  .title h3 {
    font-size: 22px;
  }

  .title h4 {
    font-size: 16px;
  }
}

.copyright {
  margin-top: 120px;
  text-align: center;
  color: #555;
  font-size: 10px;
}

.navbar {
  display: flex;
  justify-content: space-around;
  align-items: center;
  position: fixed;
  bottom: 0;
  width: 100%;
  background-color: #fff;
  padding: 10px 0;
  box-shadow: 0 -2px 5px rgba(0, 0, 0, 0.1);
}

.navbar nuxt-link {
  text-decoration: none;
  color: #000;
  font-weight: bold;
  padding: 5px 10px;
  display: flex;
  align-items: center;
}

.navbar nuxt-link i {
  margin-right: 8px;
}

.navbar nuxt-link:hover {
  color: #007bff;
}

.home {
  font-weight: normal;
  background-color: #f0f0f0;
  border-radius: 8px;
  padding: 5px 10px;
}

.home nuxt-link {
  color: #000;
}

.home nuxt-link:hover {
  color: #fff;
  background-color: #007bff;
}

@media (max-width: 768px) {
  .navbar {
    flex-direction: row;
    padding: 8px 0;
  }
}
</style>
