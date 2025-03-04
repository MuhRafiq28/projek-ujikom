<template>
  <nav class="navbar">
    <div class="logo">
      <nuxt-link class="fon" to="/">TooPiket</nuxt-link>
    </div>
    <ul class="menu" :class="{ 'menu-open': menuOpen }">
      <li class="dropdown" v-for="dropdown in dropdowns" :key="dropdown.name">
        <button class="menu-item dropdown-toggle" @click.stop="toggleDropdown(dropdown.name)">
          <i :class="dropdown.icon"></i> {{ dropdown.label }}
        </button>
        <ul :class="{ show: activeDropdown === dropdown.name }" class="dropdown-menu">
          <li v-for="item in dropdown.items" :key="item.name">
            <nuxt-link :to="item.link" class="menu-item" @click.stop="closeMenu">{{ item.name }}</nuxt-link>
          </li>
        </ul>
      </li>
      <li v-for="menu in menus" :key="menu.name">
        <nuxt-link :to="menu.link" class="menu-item" @click="closeMenu">
          <i :class="menu.icon"></i> {{ menu.name }}
        </nuxt-link>
      </li>
      <li>
        <nuxt-link to="/profile" class="menu-item" @click="closeMenu">
          <i class="fas fa-user"></i> Profil
        </nuxt-link>
      </li>
      <li>
        <nuxt-link to="/register" class="menu-item" @click="closeMenu">
          <i class="fas fa-list"></i> Register
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
      activeDropdown: null,
      dropdowns: [
        {
          name: "absensiDropdown",
          label: "Absensi",
          icon: "fas fa-clipboard-list",
          items: [
            { name: "Absensi Siswa", link: "/absensi" },
            { name: "Rekap Absensi", link: "/rekapAbsen" },
          ],
        },
        {
          name: "izinDropdown",
          label: "Siswa",
          icon: "fas fa-clipboard-check",
          items: [
            { name: "Buat Izin", link: "/ijin-siswa" },
            { name: "Daftar Izin", link: "/daftar-ijin" },
            { name: "Daftar Siswa", link: "/siswa/siswa" },
          ],
        },
        {
          name: "guruDropdown",
          label: "Guru",
          icon: "fas fa-user-tie",
          items: [
            { name: "Daftar Guru", link: "/daftarGuru" },
            { name: "Jadwal Guru", link: "/jadwalGuru" },
          ],
        },
      ],
      menus: [
        { name: "Dashboard", link: "/dashboard", icon: "fas fa-home" },
      ],
    };
  },
  async mounted() {
    await this.$store.dispatch("fetchUserRole");
  },
  computed: {
    ...mapState({
      userRole: (state) => state.user.role,
    }),
  },
  methods: {
    toggleMenu() {
      this.menuOpen = !this.menuOpen;
    },
    toggleDropdown(dropdown) {
      this.activeDropdown = this.activeDropdown === dropdown ? null : dropdown;
    },
    closeMenu() {
      this.menuOpen = false;
      this.activeDropdown = null;
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

.dropdown {
  position: relative;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  background-color: #fff;
  border-radius: 5px;
  list-style: none;
  padding: 10px 0;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  display: none;
  min-width: 180px;
}

.dropdown-menu.show {
  display: block;
}

.dropdown-menu li {
  padding: 10px 20px;
  white-space: nowrap;
}

.dropdown-toggle {
  background: none;
  border: none;
  font-size: 16px;
  color: #ecf0f1;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
}

.user-role {
  font-size: 18px;
}
</style>
