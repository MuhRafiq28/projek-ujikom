<template>
  <div>
    <Navbar />
  <div class="register-container">
    <h2 class="font-medium">Buat Akun </h2>
    <form @submit.prevent="register">
      <label>Username:</label>
      <input type="text" v-model="username" required />

      <label>Password:</label>
      <input type="password" v-model="password" required />

      <button type="submit">Buat</button>
    </form>

    <p v-if="message" class="success-message">{{ message }}</p>
    <p v-if="error" class="error-message">Gagal membuat akun</p>
  </div>
</div>
</template>

<script>
import axios from "axios";
import Navbar from "~/components/Navbar.vue";
export default {
  components: {
    Navbar
  },
  data() {
    return {
      username: "",
      password: "",
      message: "",
      error: "",
    };
  },
  methods: {
    async register() {
      try {
        const response = await axios.post("http://localhost:8080/api/register", {
          username: this.username,
          password: this.password,
        });

        this.message = "Berhasil Membuat Akun"; // Pesan sukses
        this.error = ""; // Reset error jika sukses
        this.username = "";
        this.password = "";
      } catch (err) {
        this.error = "Gagal membuat akun";
        this.message = ""; // Reset pesan sukses jika ada error
      }
    },
  },
};
</script>

<style scoped>
.register-container {
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
  background: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  color: #333;
}

form {
  display: flex;
  flex-direction: column;
}

label {
  margin-top: 10px;
}

input {
  padding: 8px;
  margin-top: 5px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

button {
  margin-top: 15px;
  padding: 10px;
  background: linear-gradient(
    135deg,
    #3498db,
    #8e44ad
  );
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #2779bd;
}

.success-message {
  color: green;
  text-align: center;
  margin-top: 10px;
}

.error-message {
  color: red;
  text-align: center;
  margin-top: 10px;
}
</style>
