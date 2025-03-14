<template>
  <div class="login-page">
    <h1>Login to TooPiket</h1>
    <form @submit.prevent="login" class="login-form">
      <div class="input-group">
        <label for="username">Username:</label>
        <input type="text" v-model="username" id="username" required />
      </div>
      <div class="input-group">
        <label for="password">Password:</label>
        <input type="password" v-model="password" id="password" required />
      </div>
      <button type="submit" class="login-btn">Login</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'LoginPage',
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    async login() {
      try {
        console.log("Proses login dimulai...");

        // Kirim permintaan ke API backend
        const response = await this.$axios.post('/api/login', {
          username: this.username,
          password: this.password
        });

        console.log("Response dari API:", response.data);

        // Ambil data yang dikembalikan oleh backend
        const { token, id, role, username } = response.data;

        if (token) {
          // Simpan data pengguna ke localStorage
          localStorage.setItem('authToken', token);
          localStorage.setItem('userId', id);
          localStorage.setItem('userRole', role);
          localStorage.setItem('username', username);

          // Update state di Vuex
          this.$store.commit("setUserRole", role);

          console.log("User data stored:", { token, id, role, username });

          // Redirect ke halaman berdasarkan role
          if (role === 'admin') {
            this.$router.push('/homenew');
          } else if (role === 'staf') {
            this.$router.push('/home-staf');
          } else {
            this.$router.push('/homenew');
          }
        } else {
          alert('Login gagal, token tidak ditemukan.');
        }
      } catch (error) {
        console.error("Login error:", error.response?.data || error.message);
        alert('Error: ' + (error.response?.data?.error || "Gagal login"));
      }
    }
  }
};
</script>

<style scoped>
.login-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(135deg, #3498db, #8e44ad);
  color: white;
  font-family: 'Arial', sans-serif;
  padding: 20px;
}

h1 {
  font-size: 2.5em;
  font-weight: bold;
  color: white;
  margin-bottom: 40px;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.3);
  animation: fadeIn 1s ease-in-out;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
  max-width: 400px;
  width: 100%;
  background-color: #ffffff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

label {
  font-size: 1.1em;
  color: #333;
}

input {
  padding: 10px;
  font-size: 1em;
  border-radius: 5px;
  border: 1px solid #ddd;
  box-sizing: border-box;
}

button.login-btn {
  padding: 15px;
  background-color: #3498db;
  color: white;
  font-size: 1.2em;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
}

button.login-btn:hover {
  background-color: #2980b9;
  transform: scale(1.05);
}

@keyframes fadeIn {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
