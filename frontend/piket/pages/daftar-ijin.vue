<template>
  <div class="daftar-ijin">
    <Navnew />
    <div class="container">
      <h1>Daftar Izin Siswa</h1>

      <!-- Input Filter -->
      <div class="filter-row">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Cari nama siswa..."
          class="search-input"
        />
        <input
          type="month"
          v-model="searchMonth"
          class="search-month"
        />
      </div>

      <div class="content-row">
        <!-- Tabel Izin -->
        <div class="left-content">
          <table>
            <thead>
              <tr>
                <th>Nama</th>
                <th>Alasan</th>
                <th>Jurusan</th>
                <th>Kelas</th>
                <th>Status</th>
                <th>Keluar</th>
                <th>Masuk</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="izin in filteredIzins"
                :key="izin.ID"
                :class="{ duplikasi: isDuplicate(izin.nama) }"
              >
                <td>{{ izin.nama }}</td>
                <td>{{ izin.alasan }}</td>
                <td>{{ izin.jurusan }}</td>
                <td>{{ izin.kelas }}</td>
                <td
                  :class="{
                    'status-keluar': izin.status === 'Keluar',
                    'status-masuk': izin.status === 'Masuk',
                    'status-terlambat': izin.status === 'Terlambat',
                    'status-pulang': izin.status === 'Pulang Lebih Cepat',
                    'status-sudah-kembali': izin.status === 'Sudah Kembali',
                  }"
                >
                  {{ izin.status }}
                </td>
                <td>
  {{
    izin.waktu_keluar
      ? formatTanggalWaktu(izin.waktu_keluar)
      : "-"
  }}
</td>
<td>
  {{
    izin.waktu_masuk
      ? formatTanggalWaktu(izin.waktu_masuk)
      : "-"
  }}
</td>


                <td>
  <button v-if="izin.status === 'Keluar'" @click="konfirmasiMasuk(izin)">
    Konfirmasi Masuk
  </button>
  <button @click="hapusIzin(izin.id)">Hapus</button>

</td>

              </tr>
            </tbody>
          </table>
        </div>

        <!-- Ringkasan -->
        <div class="right-content">
          <h2><i class="fas fa-chart-bar"></i> <span style="font-size: 12px;">Ringkasan Siswa</span></h2>
          <div
            class="card"
            v-for="item in statistikIzin"
            :key="`${item.nama}-${item.jurusan}-${item.kelas}`"
          >
            <h3><i class="fas fa-user"></i> {{ item.nama }}</h3>
            <p style="font-size: 13px;">{{ item.jurusan }} - {{ item.kelas }}</p>
            <ul>
              <li v-for="(jumlah, status) in item.status" :key="status">
                <i class="fas fa-check-circle"></i> {{ status }}: {{ jumlah }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment-timezone";
import Navnew from "../components/Navnew.vue";

export default {
  components: { Navnew },
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
      return this.izins
        .filter((izin) => {
          const namaMatch = izin.nama.toLowerCase().includes(this.searchQuery.toLowerCase());

          if (!this.searchMonth) return namaMatch;

          const izinMonth = new Date(izin.waktu_keluar).toISOString().slice(0, 7);
          return namaMatch && izinMonth === this.searchMonth;
        })
        .sort((a, b) => new Date(b.waktu_keluar) - new Date(a.waktu_keluar));
    },
    statistikIzin() {
      const countMap = {};
      this.izins.forEach((izin) => {
        const key = `${izin.nama}|${izin.jurusan}|${izin.kelas}`;
        if (!countMap[key]) {
          countMap[key] = {
            nama: izin.nama,
            jurusan: izin.jurusan,
            kelas: izin.kelas,
            status: {},
            total: 0,
          };
        }
        countMap[key].status[izin.status] = (countMap[key].status[izin.status] || 0) + 1;
        countMap[key].total += 1;
      });
      return Object.values(countMap).filter((item) => item.total > 1);
    },
  },
  methods: {
  async fetchIzin() {
    try {
      const response = await axios.get("http://localhost:8080/api/izin");
      this.izins = response.data.data.sort((a, b) => new Date(b.waktu_keluar) - new Date(a.waktu_keluar));
    } catch (error) {
      console.error("❌ Gagal mengambil data izin:", error);
    }
  },
  async konfirmasiMasuk(izin) {
    if (!izin || !izin.id) {
      alert("ID izin tidak ditemukan");
      return;
    }

    try {
      await axios.put(`http://localhost:8080/api/izin/${izin.id}/konfirmasi`);
      alert("Berhasil dikonfirmasi masuk");
      this.fetchIzin(); // refresh data
    } catch (err) {
      alert("Gagal konfirmasi: " + (err.response?.data?.error || err.message));
    }
  },
  async hapusIzin(id) {
    try {
      await axios.delete(`http://localhost:8080/api/izin/${id}`);
      this.fetchIzin();
      this.$toast?.success("🗑️ Izin berhasil dihapus.");
    } catch (err) {
      console.error("Gagal hapus:", err);
    }
  },

  formatTanggalWaktu(waktu) {
    // Format waktu dengan zona waktu Asia/Jakarta
    return moment(waktu).tz("Asia/Bangkok").format("D/M/YYYY, HH:mm:ss");
  },

  isDuplicate(nama) {
    return this.izins.filter((i) => i.nama === nama).length > 1;
  },
},

};
</script>


<style scoped>
.daftar-ijin {
  display: flex;
  margin-top: 60px;
}

.container {
  margin: auto;
  padding: 20px;
  width: 100%;
  max-width: 1200px;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.filter-row {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  margin-bottom: 20px;
  padding-right: 150px;
  padding-left: 150px;
}

.search-input,
.search-month {
  flex: 1;
  padding: 10px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid #ccc;
  width: 50px;
}

.content-row {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.left-content {
  flex: 1;
  overflow-x: auto;
}

.right-content {
  width: 200px;
  min-width: 200px;
  background: #f9f9f9;
  padding: 15px;
  border-radius: 10px;
  height: fit-content;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 30px;
}

th,
td {
  padding: 10px;
  text-align: center;
  border-bottom: 1px solid #ccc;
}

th {
  background-color: #6d7993;
  color: white;
}

.duplikasi {
  background-color: #fff3cd;
}

.status-keluar {
  color: #dc3545;
}
.status-masuk {
  color: #28a745;
}
.status-terlambat {
  color: #ffc107;
}
.status-pulang {
  color: #007bff;
}
.status-sudah-kembali {
  color: #17a2b8;
}

button {
  padding: 6px 10px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
  margin: 2px 0;
}

.btn-konfirmasi {
  background-color: #5cb85c;
  color: white;
}

button:hover {
  opacity: 0.9;
}

.card {
  background: white;
  border-left: 5px solid #6d7993;
  padding: 10px 15px;
  border-radius: 8px;
  margin-bottom: 15px;
}

.card h3 {
  margin-bottom: 8px;
  font-weight: bold;
  font-size: 14px;
}

.card ul {
  list-style: none;
  padding: 0;
}

.card li {
  margin: 5px 0;
  font-size: 13px;
}
</style>
