<template>
  <div>
    <Navbar />
    <div class="container">
      <h1>Buat Izin</h1>

      <!-- Notifikasi -->
      <div v-if="izinBerhasil" class="notif-sukses">
        {{ izinBerhasil }}
      </div>

      <!-- Form Tambah Izin -->
      <form @submit.prevent="tambahIzin">
        <input v-model="izinBaru.nama" placeholder="Nama" required />
        <input v-model="izinBaru.alasan" placeholder="Alasan" required />
        <select v-model="izinBaru.status" required>
          <option value="Keluar">Keluar</option>
          <option value="Masuk">Masuk</option>
          <option value="Terlambat">Terlambat</option>
        </select>
        <button type="submit">Tambah Izin</button>
      </form>
    </div>
  </div>
</template>


<script>
import axios from "axios";
import Navbar from "../components/Navbar.vue";

export default {
  components: {
    Navbar,
  },
  data() {
    return {
      izins: [],
      izinBaru: {
        nama: "",
        alasan: "",
        status: "Keluar",
      },
      izinBerhasil: "", // ✅ Notifikasi sukses
    };
  },
  async mounted() {
    await this.fetchIzin();
  },
  methods: {
    async fetchIzin() {
      try {
        const response = await axios.get("http://localhost:8080/api/izin");
        this.izins = response.data.data;
      } catch (error) {
        console.error("❌ Gagal mengambil data izin:", error);
      }
    },
    async tambahIzin() {
      try {
        await axios.post("http://localhost:8080/api/izin", this.izinBaru);
        this.izinBaru = { nama: "", alasan: "", status: "Keluar" }; // Reset form
        this.izinBerhasil = "✅ Izin berhasil ditambahkan!"; // ✅ Atur notifikasi sukses

        setTimeout(() => {
          this.izinBerhasil = ""; // Sembunyikan notifikasi setelah 3 detik
        }, 3000);

        this.fetchIzin();
      } catch (error) {
        console.error("❌ Gagal menambah izin:", error);
      }
    },
  },
};
</script>


<style scoped>
body {
  font-family: Arial, sans-serif;
  background: #f4f4f4;
  color: #333;
  text-align: center;
  padding: 20px;
  margin: 0;

}

.container {
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  max-width: 500px;
  margin: 40px auto;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
  margin-top: 150px;
}

h1 {
  margin-bottom: 20px;
  font-size: 30px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 10px;
  padding: 10px;
}

input, select, button {
  width: calc(100% - 20px);
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  font-size: 14px;
  margin: 0 auto;
}

button {
  background: linear-gradient(
    135deg,
    #3498db,
    #8e44ad
  );
  color: white;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.3s;
}

button:hover {
  background: linear-gradient(
    135deg,
    #0e6aa8,
    #601082
  );
}

table {
  width: calc(100% - 40px);
  margin: 20px auto;
  border-collapse: collapse;
}

th, td {
  padding: 12px;
  border-bottom: 1px solid #ddd;
  text-align: left;
}

th {
  background: #f8f8f8;
  font-weight: bold;
}

.status-keluar {
  color: #d9534f;
}

.status-masuk {
  color: #5cb85c;
}

.status-terlambat {
  color: #f0ad4e;
}
</style>
