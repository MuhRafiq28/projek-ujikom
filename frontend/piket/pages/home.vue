<template>
  <div>
    <Navbar />

    <div class="container">
      <div class="title">
        <h1>SMKN 1 Cisarua</h1>
        <h3>Halo Pa/Bu, apa yang Anda butuhkan ada di navbar. Silakan klik menu-nya.</h3>
        <p class="instruction">
          Saat membuat izin, harap lebih teliti dan pastikan siswa tidak berbohong serta memiliki alasan yang masuk akal.
        </p>
       
      </div>
      <div class="image">
        <img src="@/images/smkn.png" alt="SMKN 1 Cisarua">
      </div>
    </div>

    <footer class="footer">
      <div class="footer-content">
        <h1>TOO PIKET</h1>
        <div class="footer-section">
          <h3>SMKN 1 Cisarua</h3>
          <p>Jl. Raya Cisarua No. 1</p>
          <p>Telepon: 021-879-123-123</p>
        </div>
        <div class="footer-section">
          <h3>Jam Operasional</h3>
          <p>Senin - Jumat: 07.00 - 16.00</p>
          <p>Sabtu - Minggu: Tutup</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import Navbar from "../components/Navbar.vue";

export default {
  async mounted() {
    await this.$store.dispatch("fetchUserRole");
  },
  components: { Navbar },
  name: "HomePage",
  data() {
    return {
      userRole: "user",
    };
  },
  mounted() {
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("authToken");
      if (!token) {
        this.$router.push("/login");
      }
      this.userRole = localStorage.getItem("userRole") || "user";
      window.addEventListener("storage", this.updateUserRole);
    }
  },
  methods: {
    updateUserRole() {
      this.userRole = localStorage.getItem("userRole");
    },
  },
  beforeDestroy() {
    if (typeof window !== "undefined") {
      window.removeEventListener("storage", this.updateUserRole);
    }
  },
};
</script>

<style scoped>
.container {
  margin: 40px auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1000px;
  padding: 20px;
  margin-top: 90px;
  flex-wrap: wrap;
}

.title {
  width: 100%;
  text-align: center;
}

.title h1 {
  font-size: 36px;
  font-weight: bold;
  color: #2c3e50;
}

.title h3 {
  font-size: 18px;
  font-weight: bold;
  color: #34495e;
}

.instruction {
  font-size: 16px;
  color: #555;
  margin-top: 10px;
}

.btn-ijin {
  padding: 12px 24px;
  font-size: 16px;
  background-color: #3498db;
  color: white;
  border-radius: 8px;
  transition: all 0.3s ease-in-out;
  display: inline-block;
  margin-top: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  border: none;
  cursor: pointer;
}

.btn-ijin:hover {
  background-color: #2980b9;
  transform: scale(1.05);
}

.image {
  width: 100%;
  display: flex;
  justify-content: center;
}

.image img {
  width: 80%;
  max-width: 280px;
  border-radius: 10px;
}

.footer {
  background: linear-gradient(135deg, #3498db, #8e44ad);
  padding: 30px 0;
  color: white;
  text-align: center;
  margin-top: 50px;
}

.footer-content {
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
  max-width: 1000px;
  margin: auto;
}

.footer h1 {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 10px;
}

.footer-section h3 {
  font-size: 16px;
  font-weight: bold;
}

.footer-section p {
  font-size: 14px;
}

@media (max-width: 768px) {
  .container {
    flex-direction: column;
    text-align: center;
  }

  .title h1 {
    font-size: 28px;
  }

  .title h3 {
    font-size: 16px;
  }

  .btn-ijin {
    font-size: 14px;
  }

  .image img {
    width: 60%;
  }
}
</style>
