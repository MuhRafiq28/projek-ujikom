<template>
  <nav class="navbar">
    <div class="logo">
      <nuxt-link to="/">TooPiket</nuxt-link>
    </div>
    <button class="menu-toggle" @click="toggleMenu">
      <i class="fas" :class="menuOpen ? 'fa-times' : 'fa-bars'"></i>
    </button>
    <ul class="menu" :class="{ 'menu-open': menuOpen }">
      <li>
        <nuxt-link to="/daftarGuru" class="menu-item">
          <i class="fas fa-users"></i> Daftar Guru
        </nuxt-link>
      </li>
      <!-- Dropdown Absensi -->
      <li class="dropdown">
        <button class="menu-item dropdown-toggle" @click="toggleDropdown('absensiDropdown')">
          <i class="fas fa-clipboard-list"></i> Absensi <i ></i>
        </button>
        <ul v-show="activeDropdown === 'absensiDropdown'" class="dropdown-menu">
          <li>
            <nuxt-link to="/absensi" class="menu-item">Absensi Siswa</nuxt-link>
          </li>
          <li>
            <nuxt-link to="/rekapAbsen" class="menu-item">Rekap Absensi</nuxt-link>
          </li>
        </ul>
      </li>
      <li class="dropdown">
        <button class="menu-item dropdown-toggle" @click="toggleDropdown('absensiDropdown')">
          <i class="fas fa-clipboard-check"></i> IZIN <i ></i>
        </button>
        <ul v-show="activeDropdown === 'absensiDropdown'" class="dropdown-menu">
          <li>
            <nuxt-link to="/ijin-siswa" class="menu-item">Buat Ijin</nuxt-link>
          </li>
          <li>
            <nuxt-link to="/daftar-ijin" class="menu-item">Daftar Ijin</nuxt-link>
          </li>
        </ul>
      </li>
      <li v-for="menu in menus" :key="menu.name">
        <nuxt-link :to="menu.link" class="menu-item">
          <i :class="menu.icon"></i> {{ menu.name }}
        </nuxt-link>
      </li>
      <li>
        <nuxt-link to="/profile" class="menu-item">
          <i class="fas fa-user"></i> Profil
        </nuxt-link>
      </li>
    </ul>
    <div class="user-role">
      <h3>{{ userRole }}</h3>
    </div>
  </nav>
</template>

<script>
import { mapState } from "vuex";

export default {
  data() {
    return {
      menuOpen: false,
      activeDropdown: null, // Untuk mengontrol dropdown yang aktif
    };
  },
  async mounted() {
    await this.$store.dispatch("fetchUserRole");
  },
  computed: {
    ...mapState({
      userRole: (state) => state.user.role,
    }),
    menus() {
      const baseMenus = [
        { name: "Jadwal Guru Piket", link: "/jadwalGuru", icon: "fas fa-calendar-alt" },
  
      ];

      if (this.userRole === "admin") {
        baseMenus.push(
          { name: "Register", link: "/register", icon: "fas fa-user-plus" },
          { name: "Home", link: "/homeAdmin", icon: "fas fa-home" }
        );
      }

    },
  },
  methods: {
    toggleMenu() {
      this.menuOpen = !this.menuOpen;
    },
    toggleDropdown(dropdown) {
      // Toggle antara dropdown yang aktif
      if (this.activeDropdown === dropdown) {
        this.activeDropdown = null;
      } else {
        this.activeDropdown = dropdown;
      }
    },
  },
};
</script>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(135deg, #c60a9d, #01a0b5);
  color: white;
  padding: 15px 30px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.logo {
  font-size: 24px;
  font-weight: bold;
  color: #ecf0f1;
  text-transform: uppercase;
  letter-spacing: 2px;
  text-decoration: none;
}

.menu {
  list-style: none;
  display: flex;
  gap: 30px;
  margin: 0;
}

.menu-item {
  font-size: 16px;
  color: #ecf0f1;
  text-decoration: none;
  transition: color 0.3s ease, transform 0.3s ease;
  display: flex;
  align-items: center;
  gap: 8px;
}

.menu-item:hover {
  color: #3498db;
  transform: translateY(-2px);
}

.menu-toggle {
  display: none;
  background: none;
  border: none;
  color: white;
  font-size: 24px;
  cursor: pointer;
}

@media (max-width: 768px) {
  .menu-toggle {
    display: block;
  }

  .menu {
    position: absolute;
    top: 60px;
    left: 0;
    width: 100%;
    background: rgba(0, 0, 0, 0.8);
    flex-direction: column;
    align-items: center;
    gap: 15px;
    padding: 20px 0;
    display: none;
  }

  .menu-open {
    display: flex;
  }
}

.user-role {
  font-size: 18px;
}

/* Dropdown Menu */
.dropdown {
  position: relative;
}

.dropdown-toggle {
  background: none;
  border: none;
  color: white;
  font-size: 16px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  background-color: rgba(0, 0, 0, 0.9);
  border-radius: 5px;
  list-style: none;
  padding: 10px 0;
  display: none;
  flex-direction: column;
}

.dropdown-menu .menu-item {
  padding: 10px 20px;
  color: white;
  text-align: left;
}

.dropdown-menu .menu-item:hover {
  color: #3498db;
  transform: none;
}

.dropdown:hover .dropdown-menu {
  display: flex;
}

</style>
