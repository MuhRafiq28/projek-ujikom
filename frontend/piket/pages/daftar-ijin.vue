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
                <td>
                  <span
                    class="status-badge"
                    :class="{
                      'status-keluar': izin.status === 'Keluar',
                      'status-masuk': izin.status === 'Masuk',
                      'status-terlambat': izin.status === 'Terlambat',
                      'status-pulang': izin.status === 'Pulang Lebih Cepat',
                      'status-sudah-kembali': izin.status === 'Sudah Kembali',
                    }"
                  >
                    {{ izin.status }}
                  </span>
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
                  <button @click="konfirmasiHapus(izin)">Hapus</button>
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

    <!-- Modal Konfirmasi Masuk -->
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

    <!-- Modal Konfirmasi Hapus -->
    <div v-if="showDeleteModal" class="modal-overlay">
      <div class="modal-content">
        <h3>Konfirmasi Hapus Izin</h3>
        <p>
          Anda akan menghapus data izin untuk:
          <strong>{{ selectedIzin?.nama }}</strong> ({{ selectedIzin?.alasan }})
        </p>
        <p class="warning-text">
          <i class="fas fa-exclamation-triangle"></i> Aksi ini tidak dapat dibatalkan!
        </p>
        <div class="modal-buttons">
          <button class="delete-button" @click="prosesHapusIzin">Ya, Hapus</button>
          <button @click="batalHapus">Batal</button>
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
      showDeleteModal: false,
      selectedIzin: null,
      searchRingkasanNama: "",
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
    konfirmasiHapus(izin) {
      this.selectedIzin = izin;
      this.showDeleteModal = true;
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
    batalHapus() {
      this.showDeleteModal = false;
      this.selectedIzin = null;
    },
    async prosesHapusIzin() {
      if (!this.selectedIzin || !this.selectedIzin.id) {
        this.$toast.error("ID izin tidak ditemukan");
        return;
      }

      try {
        await axios.delete(`http://localhost:8080/api/izin/${this.selectedIzin.id}`);
        this.$toast.success("Izin berhasil dihapus");
        await this.fetchIzin();
      } catch (err) {
        this.$toast.error("Gagal menghapus izin: " + (err.response?.data?.error || err.message));
      } finally {
        this.showDeleteModal = false;
        this.selectedIzin = null;
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
/* === Global Styles === */
.daftar-ijin {
  margin-top: 70px;
  background: #f8fafc;
  min-height: 100vh;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 30px 20px;
}

h1 {
  text-align: center;
  margin-bottom: 30px;
  color: #2c3e50;
  font-weight: 600;
  font-size: 1.8rem;
}

/* === Filter Section === */
.filter-row {
  display: flex;
  gap: 15px;
  margin-bottom: 25px;
  justify-content: center;
}

.search-input, .search-month {
  padding: 12px 15px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 0.95rem;
  background: white;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  transition: all 0.2s ease;
  width: 100%;
  max-width: 300px;
}

.search-input:focus, .search-month:focus {
  outline: none;
  border-color: #6366f1;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.1);
}

/* === Main Content Layout === */
.content-row {
  display: flex;
  gap: 25px;
}

.left-content {
  flex: 1;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  padding: 20px;
}

.right-content {
  width: 280px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  padding: 20px;
  position: sticky;
  top: 90px;
  height: fit-content;
}

.right-content h2 {
  font-size: 1.2rem;
  color: #2c3e50;
  margin-bottom: 15px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.right-content h2 i {
  color: #6366f1;
}

/* === Table Styles === */
table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}

th, td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #f1f5f9;
  vertical-align: middle;
  word-wrap: break-word;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* Column Width Distribution */
th:nth-child(1), td:nth-child(1) { width: 10%; } /* Nama */
th:nth-child(2), td:nth-child(2) { width: 10%; } /* Alasan */
th:nth-child(3), td:nth-child(3) { width: 10%; } /* Jurusan */
th:nth-child(4), td:nth-child(4) { width: 8%; }  /* Kelas */
th:nth-child(5), td:nth-child(5) { width: 16%; } /* Status */
th:nth-child(6), td:nth-child(6) { width: 10%; } /* Keluar */
th:nth-child(7), td:nth-child(7) { width: 10%; } /* Masuk */
th:nth-child(8), td:nth-child(8) { width: 8%; }  /* Aksi */

th {
  background: #4f46e5;
  color: white;
  font-weight: 500;
  text-transform: uppercase;
  font-size: 0.8rem;
  letter-spacing: 0.5px;
  position: sticky;
  top: 0;
}

td {
  font-size: 0.9rem;
  color: #334155;
}

/* Status Badges */
.status-badge {
  display: inline-block;
  padding: 4px 10px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 500;
  text-align: center;
  min-width: 100px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.status-keluar {
  background: #fee2e2;
  color: #dc2626;
  border: 1px solid #fecaca;
}
.status-masuk {
  background: #dcfce7;
  color: #16a34a;
  border: 1px solid #bbf7d0;
}
.status-terlambat {
  background: #fef9c3;
  color: #ca8a04;
  border: 1px solid #fef08a;
}
.status-pulang {
  background: #dbeafe;
  color: #2563eb;
  border: 1px solid #bfdbfe;
}
.status-sudah-kembali {
  background: #cffafe;
  color: #0891b2;
  border: 1px solid #a5f3fc;
}

/* Time Columns */
td:nth-child(6), td:nth-child(7) {
  font-family: monospace;
  font-size: 0.85rem;
  color: #475569;
}

/* Action Buttons */
button {
  padding: 8px 12px;
  border: none;
  border-radius: 6px;
  font-size: 0.85rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  margin-right: 5px;
}

button:first-of-type {
  background: #4f46e5;
  color: white;
}

button:last-of-type {
  background: #ef4444;
  color: white;
}

button:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

/* Card Styles */
.card {
  background: white;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 15px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
  border-left: 3px solid #6366f1;
}

.card h3 {
  font-size: 0.95rem;
  color: #2c3e50;
  margin-bottom: 5px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.card h3 i {
  color: #6366f1;
  font-size: 0.9rem;
}

.card p {
  font-size: 0.8rem;
  color: #64748b;
  margin-bottom: 10px;
}

.card ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.card li {
  font-size: 0.8rem;
  color: #475569;
  padding: 5px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.card li i {
  color: #10b981;
  font-size: 0.7rem;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 25px;
  border-radius: 12px;
  width: 90%;
  max-width: 400px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  animation: modalFadeIn 0.2s ease-out;
}

.modal-content h3 {
  color: #2c3e50;
  margin-bottom: 15px;
  font-size: 1.2rem;
}

.modal-content p {
  color: #64748b;
  margin-bottom: 20px;
  line-height: 1.5;
}

.warning-text {
  color: #dc2626;
  font-weight: 500;
  margin: 15px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.warning-text i {
  font-size: 1rem;
}

.modal-content p strong {
  color: #2c3e50;
}

.modal-buttons {
  display: flex;
  gap: 10px;
  justify-content: center;
}

.modal-buttons button {
  flex: 1;
  padding: 10px;
  margin: 0;
}

.modal-buttons button:first-child {
  background: #4f46e5;
}

.delete-button {
  background: #ef4444 !important;
}

.delete-button:hover {
  background: #dc2626 !important;
}

.modal-buttons button:last-child {
  background: #64748b;
}

@keyframes modalFadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* === Responsive Design === */
@media (max-width: 1200px) {
  th:nth-child(2), td:nth-child(2) { width: 18%; }
  th:nth-child(6), td:nth-child(6),
  th:nth-child(7), td:nth-child(7) { width: 10%; }
}

@media (max-width: 992px) {
  .content-row {
    flex-direction: column;
  }

  .right-content {
    width: 100%;
    position: static;
  }

  th:nth-child(1), td:nth-child(1) { width: 20%; }
  th:nth-child(2), td:nth-child(2) { width: 25%; }
  th:nth-child(5), td:nth-child(5) { width: 12%; }
}

@media (max-width: 768px) {
  .filter-row {
    flex-direction: column;
    align-items: stretch;
  }

  .search-input, .search-month {
    max-width: 100%;
  }

  /* Hide less important columns */
  th:nth-child(3), td:nth-child(3), /* Jurusan */
  th:nth-child(4), td:nth-child(4) { /* Kelas */
    display: none;
  }

  /* Adjust remaining columns */
  th:nth-child(1), td:nth-child(1) { width: 25%; }
  th:nth-child(2), td:nth-child(2) { width: 30%; }
  th:nth-child(5), td:nth-child(5) { width: 20%; }
  th:nth-child(6), td:nth-child(6),
  th:nth-child(7), td:nth-child(7) { width: 12%; }
}

@media (max-width: 576px) {
  /* Hide time columns on very small screens */
  th:nth-child(6), td:nth-child(6),
  th:nth-child(7), td:nth-child(7) {
    display: none;
  }

  /* Final column adjustment */
  th:nth-child(1), td:nth-child(1) { width: 30%; }
  th:nth-child(2), td:nth-child(2) { width: 40%; }
  th:nth-child(5), td:nth-child(5) { width: 30%; }

  /* Make buttons smaller */
  button {
    padding: 6px 8px;
    font-size: 0.75rem;
  }
}
</style>
