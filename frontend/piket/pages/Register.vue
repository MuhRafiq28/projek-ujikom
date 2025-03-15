<template>
  <div>

    <Navnew/>

    <div class="container">
      <div class="card">
        <div class="form-container">
          <h2 class="font-medium">Buat Akun</h2>
          <form @submit.prevent="register">
            <label>Username:</label>
            <input type="text" v-model="username" required />

            <label>Password:</label>
            <input type="password" v-model="password" required />

            <button type="submit">Buat</button>
          </form>
        </div>
        <div class="image-container">
          <img src="/images/inspirasi-logo.png" alt="Image" />
        </div>
      </div>

      <!-- Modal Success -->
      <div v-if="showSuccessModal" class="modal">
        <div class="modal-content">
          <h3>Berhasil Membuat Akun</h3>
          <p>{{ message }}</p>
          <button @click="closeModal">Tutup</button>
        </div>
      </div>

      <!-- Modal Error -->
      <div v-if="showErrorModal" class="modal">
        <div class="modal-content">
          <h3>Gagal Membuat Akun</h3>
          <p>{{ error }}</p>
          <button @click="closeModal">Tutup</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Navnew from "~/components/Navnew.vue";


export default {
  components: {
    Navnew,
  },
  data() {
    return {
      username: "",
      password: "",
      message: "",
      error: "",
      showSuccessModal: false,
      showErrorModal: false,
      userRole: "user", // Default userRole
    };
  },
  mounted() {
    // Ambil peran user dari localStorage
    this.userRole = localStorage.getItem("userRole") || "user";
  },
  methods: {
    async register() {
      try {
        const response = await axios.post("http://localhost:8080/api/register", {
          username: this.username,
          password: this.password,
        });

        this.$toast.success("Akun berhasil dibuat!", {
      position: "top-right",
      timeout: 3000
    });
        this.error = "";
        this.showSuccessModal = true;
        this.showErrorModal = false;
        this.username = "";
        this.password = "";
      } catch (err) {
        this.error = "Gagal membuat akun";
        this.message = "";
        this.showErrorModal = true;
        this.showSuccessModal = false;
      }
    },
    closeModal() {
      this.showSuccessModal = false;
      this.showErrorModal = false;
    },
  },
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 140px;

}

.card {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  max-width: 700px;
  width: 100%;
  background: #fff;
  border-radius: 15px;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
  padding: 20px;
}

.form-container {
  width: 50%;
  padding-right: 20px;
}

h2 {
  text-align: center;
  color: #333;
  font-family: 'Poppins', sans-serif;
  margin-bottom: 20px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

label {
  font-size: 14px;
  color: #555;
}

input {
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.3s ease;
}

input:focus {
  border-color: #3498db;
  box-shadow: 0 0 5px rgba(52, 152, 219, 0.5);
}

input:hover {
  border-color: #8e44ad;
  background-color: #f4f4f9;
  transform: scale(1.05); /* Membesar input saat hover */
}

button {
  padding: 12px;
  background: linear-gradient(135deg, #3498db, #8e44ad);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.3s ease;
}

button:hover {
  background-color: #2779bd;
  transform: translateY(-2px);
}

button:active {
  transform: translateY(2px);
}

.image-container {
  width: 50%;
  padding-left: 20px;
}

.image-container img {
  width: 100%;
  height: auto;
  max-height: 100%;
  object-fit: cover;
  border-radius: 10px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: #fff;
  padding: 30px;
  border-radius: 10px;
  text-align: center;
  width: 350px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.modal button {
  padding: 12px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.modal button:hover {
  background-color: #2779bd;
  transform: translateY(-2px);
}

@keyframes slideIn {
  from {
    transform: translateY(-20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
