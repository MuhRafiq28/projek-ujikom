<template>
  <div>
    <Navnew />
    <div class="container">
      <h1>Daftar Izin Siswa</h1>

      <!-- Input Pencarian -->
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Cari berdasarkan nama..."
        class="search-input"
      />

      <!-- Input Filter Bulan -->
      <input
        type="month"
        v-model="searchMonth"
        class="search-month"
      />

      <!-- Tabel Izin -->
      <table>
        <thead>
          <tr>
            <th>Nama</th>
            <th>Alasan</th>
            <th>Status</th>
            <th>Waktu Keluar</th>
            <th>Waktu Masuk</th>
            <th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="izin in filteredIzins"
            :key="izin.ID"
            :class="{ 'duplikasi': isDuplicate(izin.nama) }"
          >
            <td>{{ izin.nama }}</td>
            <td>{{ izin.alasan }}</td>
            <td
              :class="{
                'status-keluar': izin.status === 'Keluar',
                'status-masuk': izin.status === 'Masuk',
                'status-terlambat': izin.status === 'Terlambat',
              }"
            >
              {{ izin.status }}
            </td>
            <td>{{ formatTanggalWaktu(izin.waktu_keluar) }}</td>
            <td>{{ izin.waktu_masuk ? formatTanggalWaktu(izin.waktu_masuk) : "-" }}</td>
            <td>
              <button class="btn-konfirmasi"
                v-if="izin.status === 'Keluar'"
                @click="konfirmasiMasuk(izin.ID)"
              >
                Konfirmasi Masuk
              </button>
              <button @click="hapusIzin(izin.ID)">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Navnew from "../components/Navnew.vue";

export default {
  components: {
    Navnew,
  },
  data() {
    return {
      izins: [],
      searchQuery: "",
      searchMonth: "",
    };
  },
  async mounted() {
    await this.fetchIzin();
  },
  computed: {
    filteredIzins() {
      return this.izins.filter((izin) => {
        const namaMatch = izin.nama.toLowerCase().includes(this.searchQuery.toLowerCase());

        if (!this.searchMonth) return namaMatch;

        const izinDate = new Date(izin.waktu_keluar);
        const izinMonth = izinDate.toISOString().slice(0, 7); // Format YYYY-MM

        return namaMatch && izinMonth === this.searchMonth;
      });
    },
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
    async konfirmasiMasuk(id) {
      if (!id) {
        console.error("❌ ID izin tidak ditemukan di frontend!");
        return;
      }
      try {
        await axios.put(`http://localhost:8080/api/izin/${id}/konfirmasi`);
        this.fetchIzin();
      } catch (error) {
        console.error(
          "❌ Gagal mengonfirmasi masuk:",
          error.response?.data || error
        );
      }
    },
    async hapusIzin(id) {
      if (!id) {
        console.error("❌ ID izin tidak ditemukan di frontend!");
        return;
      }
      try {
        await axios.delete(`http://localhost:8080/api/izin/${id}`);
        this.fetchIzin();
      } catch (error) {
        console.error(
          "❌ Gagal menghapus izin:",
          error.response?.data || error
        );
      }
    },
    formatTanggalWaktu(waktu) {
      if (!waktu) return "-";
      const date = new Date(waktu);
      return `${date.toLocaleDateString("id-ID", {
        weekday: "long",
        day: "2-digit",
        month: "long",
        year: "numeric",
      })}, ${date.toLocaleTimeString("id-ID", {
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
      })}`;
    },
    isDuplicate(nama) {
      return this.izins.filter((izin) => izin.nama === nama).length > 1;
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
  margin: 0 auto;
  background: white;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  margin-top: 70px;
  text-align: center;
}

h1 {
  margin-bottom: 20px;
  font-size: 24px;
}

.search-input,
.search-month {
  width: 100%;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #ccc;
  font-size: 16px;
  margin: 5px 0;
  display: block;
  text-align: center;
}

button {
  background: linear-gradient(135deg, #9cc7e4, #6D7993);
  color: white;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.3s;
  border: none;
  padding: 10px 15px;
  border-radius: 8px;
}

button:hover {
  background: #555;
}

table {
  width: 100%;
  margin: 20px 0;
  border-collapse: collapse;
  border-radius: 10px;
  overflow: hidden;
}

th,
td {
  padding: 12px;
  border-bottom: 1px solid #ddd;
  text-align: left;
}

th {
  background-color: #96858F;
  font-weight: bold;
  color: white;
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

.duplikasi {
  background-color: #fff3cd !important;
}
</style>
