<template>
  <div class="profile">
    <Navbar />
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
      <button @click="logout" class="logout-btn">Logout</button> <!-- Tombol logout -->
    </div>
  </div>
</template>

<script>
import Navbar from '~/components/Navbar.vue';

export default {
  components: { Navbar },
  data() {
    return {
      profile: {
        username: '',
        email: '',
        role: '',
      },
    };
  },
  mounted() {
    // Memastikan kode ini hanya dijalankan di sisi klien
    if (typeof window !== 'undefined' && window.localStorage) {
      this.profile.username = localStorage.getItem('username') || '';
      this.profile.email = localStorage.getItem('userEmail') || '';
      this.profile.role = localStorage.getItem('userRole') || '';
    }
  },
  methods: {
    saveProfile() {
      // Simulasi penyimpanan data baru (misalnya, kirim ke API)
      if (typeof window !== 'undefined' && window.localStorage) {
        localStorage.setItem('userEmail', this.profile.email);
      }
      alert("Profil berhasil diperbarui!");
    },
    logout() {
      console.log("Logout dipanggil"); // Debug log
      localStorage.removeItem("authToken");
      localStorage.removeItem("userId");
      localStorage.removeItem("userRole");
      localStorage.removeItem("username");
      console.log("Token setelah logout:", localStorage.getItem("authToken")); // Debug log
      this.$router.push("/landing"); // Redirect ke halaman login
    },
  },
};
</script>

<style scoped>
.profile-page {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  background-color: #f8f9fa;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

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

button.save-btn {
  padding: 10px 20px;

  background: linear-gradient(
    135deg,
    #3498db,
    #8e44ad
  );
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

button.logout-btn {
  padding: 10px 20px;

  background: linear-gradient(
    135deg,
    #b434db,
    #e10000
  );
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
</style>
