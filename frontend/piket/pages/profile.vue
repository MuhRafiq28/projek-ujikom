<template>
  <div>
    <!-- Layout Admin -->
    <div class="admin-layout" v-if="isAdmin">
      <!-- Navbar Admin -->
      <NavAdmin />
      <!-- Container Profil untuk Admin -->
      <div class="profile-container-admin">
        <div class="profile-page">
          <h1>Profil Pengguna</h1>
          <div class="profile-info">
            <div class="profile-item">
              <label for="username">Nama Pengguna:</label>
              <input v-model="profile.username" type="text" id="username" disabled />
            </div>
            <div class="profile-item">
              <label for="email">Email:</label>
              <input v-model="profile.email" type="email" id="email" />
            </div>
            <div class="profile-item">
              <label for="role">Peran:</label>
              <input v-model="profile.role" type="text" id="role" disabled />
            </div>
          </div>
          <button @click="saveProfile" class="save-btn">Simpan Perubahan</button>
          <button @click="logout" class="logout-btn">Logout</button>
        </div>
      </div>
    </div>

    <!-- Layout User (Navbar biasa) -->
    <div v-else>
      <Navbar />
      <div class="profile-container">
        <div class="profile-page">
          <h1>Profil Pengguna</h1>
          <div class="profile-info">
            <div class="profile-item">
              <label for="username">Nama Pengguna:</label>
              <input v-model="profile.username" type="text" id="username" disabled />
            </div>
            <div class="profile-item">
              <label for="email">Email:</label>
              <input v-model="profile.email" type="email" id="email" />
            </div>
            <div class="profile-item">
              <label for="role">Peran:</label>
              <input v-model="profile.role" type="text" id="role" disabled />
            </div>
          </div>
          <button @click="saveProfile" class="save-btn">Simpan Perubahan</button>
          <button @click="logout" class="logout-btn">Logout</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from '~/components/Navbar.vue';
import NavAdmin from '~/components/NavAdmin.vue';

export default {
  components: {
    Navbar,
    NavAdmin
  },
  data() {
    return {
      profile: {
        username: '',
        email: '',
        role: '',
      },
    };
  },
  computed: {
    isAdmin() {
      if (process.client) {
        return localStorage.getItem('userRole') === 'admin';
      }
      return false;
    }
  },
  mounted() {
    if (process.client) {
      this.profile.username = localStorage.getItem('username') || '';
      this.profile.email = localStorage.getItem('userEmail') || '';
      this.profile.role = localStorage.getItem('userRole') || '';
    }
  },
  methods: {
    saveProfile() {
      if (process.client) {
        localStorage.setItem('userEmail', this.profile.email);
      }
      alert("Profil berhasil diperbarui!");
    },
    logout() {
      if (process.client) {
        localStorage.removeItem("authToken");
        localStorage.removeItem("userId");
        localStorage.removeItem("userRole");
        localStorage.removeItem("username");
      }
      this.$router.push("/landing");
    },
  },
};
</script>

<style scoped>
/* Layout utama dengan flexbox */
.admin-layout {
  display: flex;
  height: 100vh;
}

/* Profile Container untuk Admin Layout */
.profile-container-admin {
  flex: 1;
  padding: 40px;

  display: flex;

  align-items: center;
  margin-left: 200px; /* Memberikan jarak agar profil tidak mepet navbar */
}

.profile-page {
  width: 100%;
  max-width: 600px;
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

/* Styling teks dan input */
h1 {
  font-size: 2em;
  margin-bottom: 20px;
  text-align: center;
}

.profile-info {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.profile-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

label {
  font-size: 1em;
}

input {
  width: 100%;
  padding: 8px;
  font-size: 1em;
  border-radius: 5px;
  border: 1px solid #ddd;
  box-sizing: border-box;
}

/* Tombol simpan */
button.save-btn {
  padding: 10px 20px;
  background: linear-gradient(135deg, #3498db, #8e44ad);
  color: white;
  border: none;
  border-radius: 5px;
  width: 100%;
  font-size: 1.1em;
  cursor: pointer;
  margin-top: 20px;
}

button.save-btn:hover {
  background-color: #2980b9;
}

/* Tombol logout */
button.logout-btn {
  padding: 10px 20px;
  background: linear-gradient(135deg, #b434db, #e10000);
  color: white;
  border: none;
  border-radius: 5px;
  width: 100%;
  font-size: 1.1em;
  cursor: pointer;
  margin-top: 10px;
}

button.logout-btn:hover {
  background-color: #c0392b;
}

/* Profile Container untuk User Layout */
.profile-container {
  margin-top: 80px;
  flex: 1;
  padding: 40px;
  background-color: #f8f9fa;
  display: flex;
  justify-content: center; /* Menjaga konten berada di tengah */
  align-items: center;
}
</style>
