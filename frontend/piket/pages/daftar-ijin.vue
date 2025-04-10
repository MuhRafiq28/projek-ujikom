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
        <input type="month" v-model="searchMonth" class="search-month" />
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
              <tr v-for="izin in filteredIzins" :key="izin.ID">
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
                  <button
                    v-if="izin.status === 'Keluar'"
                    @click="konfirmasiMasuk(izin)"
                  >
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
          <h2>
            <i class="fas fa-chart-bar"></i>
            <span style="font-size: 12px">Ringkasan Siswa</span>
          </h2>

          <input
            v-model="searchRingkasanNama"
            placeholder="Cari nama "
            class="search-input"
            style="margin-bottom: 10px; font-size: 13px; width: 140px;"
          />

          <div
            class="card"
            v-for="item in filteredStatistikIzin"
            :key="`${item.nama}-${item.jurusan}-${item.kelas}`"
          >
            <h3><i class="fas fa-user"></i> {{ item.nama }}</h3>
            <p style="font-size: 13px">{{ item.jurusan }} - {{ item.kelas }}</p>
            <ul>
              <li v-for="(jumlah, status) in item.status" :key="status">
                <i class="fas fa-check-circle"></i> {{ status }}: {{ jumlah }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Konfirmasi -->
    <div v-if="showModal" class="modal-overlay">
      <div class="modal-content">
        <h3>Konfirmasi Masuk</h3>
        <p>
          Yakin ingin konfirmasi bahwa
          <strong>{{ selectedIzin?.nama }}</strong> sudah masuk?
        </p>
        <div class="modal-buttons">
          <button @click="prosesKonfirmasiMasuk">Ya, Konfirmasi</button>
          <button @click="batalKonfirmasi">Batal</button>
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
      showModal: false,
      selectedIzin: null,
      searchRingkasanNama: "", // 👈 tambahan untuk filter ringkasan
    };
  },
  async mounted() {
    await this.fetchIzin();
  },
  computed: {
    filteredIzins() {
      return this.izins
        .filter((izin) => {
          const namaMatch = izin.nama
            .toLowerCase()
            .includes(this.searchQuery.toLowerCase());

          if (!this.searchMonth) return namaMatch;

          const izinMonth = new Date(izin.waktu_keluar)
            .toISOString()
            .slice(0, 7);
          return namaMatch && izinMonth === this.searchMonth;
        })
        .sort((a, b) => new Date(b.waktu_keluar) - new Date(a.waktu_keluar));
    },
    statistikIzin() {
      const statistik = [];
      const defaultStatusList = {
        'Keluar': 0,
        'Masuk': 0,
        'Terlambat': 0,
        'Pulang Lebih Cepat': 0,
        'Sudah Kembali': 0,
      };

      this.filteredIzins.forEach((izin) => {
        const key = `${izin.nama}-${izin.jurusan}-${izin.kelas}`;
        const existing = statistik.find(item => item.key === key);
        const normalizedStatus = izin.status.trim();

        if (existing) {
          if (!existing.status[normalizedStatus]) {
            existing.status[normalizedStatus] = 0;
          }
          existing.status[normalizedStatus] += 1;
          existing.count += 1;
        } else {
          const statusObj = { ...defaultStatusList };
          if (statusObj[normalizedStatus] !== undefined) {
            statusObj[normalizedStatus] = 1;
          } else {
            statusObj[normalizedStatus] = 1;
          }

          statistik.push({
            key,
            nama: izin.nama,
            jurusan: izin.jurusan,
            kelas: izin.kelas,
            status: statusObj,
            count: 1,
          });
        }
      });

      return statistik.filter(item => item.count > 1);
    },
    filteredStatistikIzin() {
      return this.statistikIzin.filter((item) =>
        item.nama.toLowerCase().includes(this.searchRingkasanNama.toLowerCase())
      );
    },
  },
  methods: {
    async fetchIzin() {
      try {
        const response = await axios.get("http://localhost:8080/api/izin");
        this.izins = response.data.data.sort(
          (a, b) => new Date(b.waktu_keluar) - new Date(a.waktu_keluar)
        );
      } catch (error) {
        console.error("❌ Gagal mengambil data izin:", error);
      }
    },
    konfirmasiMasuk(izin) {
      this.selectedIzin = izin;
      this.showModal = true;
    },
    async prosesKonfirmasiMasuk() {
      if (!this.selectedIzin || !this.selectedIzin.id) {
        alert("ID izin tidak ditemukan");
        return;
      }

      try {
        await axios.put(
          `http://localhost:8080/api/izin/${this.selectedIzin.id}/konfirmasi`
        );
        this.showModal = false;
        this.selectedIzin = null;
        await this.fetchIzin();
        this.$toast?.success("✅ Berhasil dikonfirmasi masuk");
      } catch (err) {
        alert("Gagal konfirmasi: " + (err.response?.data?.error || err.message));
      }
    },
    batalKonfirmasi() {
      this.showModal = false;
      this.selectedIzin = null;
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
      if (!waktu) return "-";
      return moment(waktu).tz("Asia/Bangkok").format("D/M/YYYY, HH:mm:ss");
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
  width: 210px;
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

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: #fff;
  padding: 25px;
  border-radius: 10px;
  width: 90%;
  max-width: 400px;
  text-align: center;
}

.modal-buttons button {
  margin: 10px;
  padding: 8px 15px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
}

.modal-buttons button:first-child {
  background-color: #28a745;
  color: white;
}

.modal-buttons button:last-child {
  background-color: #dc3545;
  color: white;
}

</style>
