<template>
  <div>
    <Navbar />
    <div class="container">
      <h1>Daftar Siswa</h1>

      <!-- Filter Jurusan -->
      <div class="filter-container">
        <button
          v-for="j in jurusanList"
          :key="j"
          class="btn filter-btn m-1"
          :class="{
            'btn-success': j === 'PH',
            'btn-primary': j === 'RPL',
            'btn-pink': j === 'MPLB'
          }"
          @click="filterJurusan = (filterJurusan === j) ? null : j">
          {{ j }}
        </button>
      </div>

      <!-- Filter Kelas -->
      <div class="filter-container">
        <button
          v-for="k in kelasList"
          :key="k"
          class="btn filter-btn m-1"
          :class="getKelasButtonColor(k)"
          @click="filterKelas = (filterKelas === k) ? null : k">
          {{ k }}
        </button>
      </div>

      <!-- Tabel Siswa -->
      <table class="table table-striped table-bordered">
        <thead>
          <tr>
            <th>ID</th>
            <th>NIS</th>
            <th>Nama</th>
            <th>Jurusan</th>
            <th>Kelas</th>
            <th>Hadir</th>
            <th>Sakit</th>
            <th>Izin</th>
            <th>Alfa</th>
          </tr>
        </thead>
        <tbody>
          <tr v-if="filteredSiswa.length === 0">
            <td colspan="9" class="text-center">Tidak ada data siswa</td>
          </tr>
          <tr v-for="siswa in filteredSiswa" :key="siswa.id">
            <td>{{ siswa.id }}</td>
            <td>{{ siswa.nis }}</td>
            <td @click="showModal(siswa)" class="nama-siswa">
              {{ siswa.nama }}
            </td>
            <td>{{ siswa.jurusan }}</td>
            <td>{{ siswa.kelas }}</td>
            <td :class="'text-success fw-bold'">{{ siswa.jumlahHadir }}</td>
            <td :class="'text-primary fw-bold'">{{ siswa.jumlahSakit }}</td>
            <td :class="'text-warning fw-bold'">{{ siswa.jumlahIzin }}</td>
            <td :class="'text-danger fw-bold'">{{ siswa.jumlahAlfa }}</td>
          </tr>
        </tbody>
      </table>

      <!-- Modal untuk menampilkan detail absensi -->
      <RekapAbsensiModal
        :show="showModalFlag"
        :siswa="selectedSiswa"
        :absensi="absensiSiswa"
        @close="showModalFlag = false"
      />
    </div>
  </div>
</template>

<script>
import RekapAbsensiModal from "@/components/RekapAbsensiModal.vue";
import Navbar from "../components/Navbar.vue";
import axios from "axios";

export default {
  components: {
    RekapAbsensiModal,
    Navbar,
  },
  data() {
    return {
      rekapAbsensi: [],
      showModalFlag: false,
      selectedSiswa: null,
      absensiSiswa: [],
      filterJurusan: null,
      filterKelas: null,
      jurusanList: ["RPL", "MPLB", "PH"],
      kelasList: ["10A", "10B", "10C", "11A", "11B", "11C", "12A", "12B", "12C"],
    };
  },
  computed: {
    filteredSiswa() {
      return this.rekapAbsensi.filter(siswa => {
        return (
          (!this.filterJurusan || siswa.jurusan === this.filterJurusan) &&
          (!this.filterKelas || siswa.kelas === this.filterKelas)
        );
      });
    },
  },
  created() {
    this.fetchAllSiswa();
  },
  methods: {
    async fetchAllSiswa() {
      try {
        const response = await axios.get("http://localhost:8080/api/siswa");
        const siswaData = response.data || [];

        const absensiPromises = siswaData.map(async (siswa) => {
          try {
            const absensiResponse = await axios.get(
              `http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(siswa.nama)}`
            );

            const absensi = (absensiResponse.data.status === "success" && Array.isArray(absensiResponse.data.data))
              ? absensiResponse.data.data
              : [];

            siswa.jumlahHadir = absensi.filter(a => a.status.trim().toLowerCase() === "hadir").length;
            siswa.jumlahSakit = absensi.filter(a => a.status.trim().toLowerCase() === "sakit").length;
            siswa.jumlahIzin = absensi.filter(a => a.status.trim().toLowerCase() === "izin").length;
            siswa.jumlahAlfa = absensi.filter(a => a.status.trim().toLowerCase() === "alfa").length;
          } catch (error) {
            siswa.jumlahHadir = 0;
            siswa.jumlahSakit = 0;
            siswa.jumlahIzin = 0;
            siswa.jumlahAlfa = 0;
          }
        });

        await Promise.all(absensiPromises);
        this.rekapAbsensi = siswaData;
      } catch (error) {
        console.error("Error fetching all siswa data:", error);
        this.rekapAbsensi = [];
      }
    },

    getKelasButtonColor(kelas) {
      if (kelas.startsWith("10")) return "btn-success"; // Hijau untuk kelas 10
      if (kelas.startsWith("11")) return "btn-warning"; // Kuning untuk kelas 11
      if (kelas.startsWith("12")) return "btn-danger";  // Merah untuk kelas 12
      return "btn-secondary";
    },

    async showModal(siswa) {
      this.selectedSiswa = siswa;
      this.showModalFlag = true;
      try {
        const response = await axios.get(`http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(siswa.nama)}`);
        this.absensiSiswa = response.data || [];
      } catch (error) {
        console.error("Error fetching absensi:", error);
        this.absensiSiswa = [];
      }
    },
  },
};
</script>

<style scoped>
.container {
  max-width: 1000px;
  margin: 0 auto;
  margin-top: 90px;
}

.filter-container {
  margin-bottom: 10px;
}

/* Efek hover tombol */
.filter-btn {
  transition: transform 0.3s ease-in-out, rotate 0.2s;
  box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.1);
}

.filter-btn:hover {
  transform: scale(1.12) rotate(3deg);
  box-shadow: 4px 4px 12px rgba(0, 0, 0, 0.3);
}

/* Hover Nama Siswa */
.nama-siswa {
  font-weight: bold;
  font-size: 1.1rem;
  color: #007bff;
  cursor: pointer;
  transition: transform 0.3s ease-in-out;
}

.nama-siswa:hover {
  color: #ff4500;
  transform: scale(1.2);
}

/* Warna absensi */
.text-success { color: #28a745 !important; }
.text-primary { color: #007bff !important; }
.text-warning { color: #ffc107 !important; }
.text-danger { color: #dc3545 !important; }
</style>
