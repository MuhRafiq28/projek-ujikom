<template>
  <nav
    class="navbar navbar-expand-lg shadow-sm fixed-top"
    :class="{ 'navbar-hidden': !isVisible }"
  >
    <div class="container">
      <NuxtLink class="navbar-brand fw-bold" to="/">
        <i class="fas fa-school"></i> Too<span style="color: #6D7993;">piket</span>
      </NuxtLink>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse justify-content-center" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="absensiDropdown" role="button" data-bs-toggle="dropdown">
              <i class="fas fa-clipboard-list"></i> Absensi
            </a>
            <ul class="dropdown-menu">
              <li><NuxtLink class="dropdown-item" to="/absensi"><i class="fas fa-user-check"></i> Absensi Siswa</NuxtLink></li>
              <li><NuxtLink class="dropdown-item" to="/rekapAbsen"><i class="fas fa-chart-bar"></i> Rekap Absensi</NuxtLink></li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="siswaDropdown" role="button" data-bs-toggle="dropdown">
              <i class="fas fa-user-graduate"></i> Siswa
            </a>
            <ul class="dropdown-menu">
              <li v-if="role === 'user'"><NuxtLink class="dropdown-item" to="/ijin-siswa"><i class="fas fa-clipboard"></i> Buat Ijin</NuxtLink></li>
              <li><NuxtLink class="dropdown-item" to="/daftar-ijin"><i class="fas fa-list"></i> Daftar Ijin</NuxtLink></li>
              <li><NuxtLink class="dropdown-item" to="/siswa/siswa"><i class="fas fa-users"></i> Daftar Siswa</NuxtLink></li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="guruDropdown" role="button" data-bs-toggle="dropdown">
              <i class="fas fa-chalkboard-teacher"></i> Guru
            </a>
            <ul class="dropdown-menu">
              <li><NuxtLink class="dropdown-item" to="/daftarGuru"><i class="fas fa-user-tie"></i> Daftar Guru</NuxtLink></li>
              <li><NuxtLink class="dropdown-item" to="/jadwalGuru"><i class="fas fa-calendar-alt"></i> Jadwal Guru</NuxtLink></li>
            </ul>
          </li>
          <li v-if="role === 'admin'" class="nav-item">
            <NuxtLink class="nav-link" to="/register"><i class="fas fa-user-plus"></i> Register</NuxtLink>
          </li>
          <li class="nav-item">
            <NuxtLink class="nav-link" to="/profile"><i class="fas fa-user"></i> Profil</NuxtLink>
          </li>
        </ul>
      </div>
      <div class="role-name">
        <span class="nav-link fw-bold" style="color: #6D7993;"><i class="fas fa-user-circle"></i> {{ username }}</span>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      role: "user",
      username: "Guest",
      isVisible: true, // Untuk menyimpan status navbar (muncul/tidak)
      lastScrollPosition: 0, // Untuk menyimpan posisi scroll terakhir
    };
  },
  mounted() {
    if (process.client) {
      this.role = localStorage.getItem("userRole") || "user";
      this.username = localStorage.getItem("username") || "Guest";

      // Tambahkan event listener untuk mendeteksi scroll
      window.addEventListener("scroll", this.handleScroll);
    }
  },
  beforeDestroy() {
    if (process.client) {
      // Hapus event listener saat komponen dihancurkan
      window.removeEventListener("scroll", this.handleScroll);
    }
  },
  methods: {
    handleScroll() {
      const currentScroll = window.pageYOffset || document.documentElement.scrollTop;

      if (currentScroll > this.lastScrollPosition) {
        // Scroll ke bawah, sembunyikan navbar
        this.isVisible = false;
      } else {
        // Scroll ke atas, tampilkan navbar
        this.isVisible = true;
      }

      this.lastScrollPosition = currentScroll; // Simpan posisi scroll terakhir
    },
  },
};
</script>

<style scoped>
.navbar {
  background: #D5D5D5;
  height: 70px;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 999; /* Lebih kecil dari modal */
  transition: transform 0.3s ease-in-out;
}

.navbar-hidden {
  transform: translateY(-100%);
}

.navbar-nav {
  text-align: center;
}

.navbar-nav .nav-item {
  margin: 0 15px;
}

.navbar-nav .nav-link {
  color: black;
  font-weight: 500;
  transition: all 0.3s ease-in-out;
  position: relative;
}

.navbar-nav .nav-link:hover {
  color: #96858F;
  transform: scale(1.05);
}

.navbar-nav .dropdown-item {
  transition: all 0.3s ease-in-out;
}

.navbar-nav .dropdown-item:hover {
  background: #96858F;
  color: white;
  transform: translateX(5px);
}

.role-name {
  margin-left: 20px;
}
</style>
