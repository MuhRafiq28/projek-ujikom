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
          class="btn btn-primary m-1"
          @click="filterJurusan = (filterJurusan === j) ? null : j">
          {{ j }}
        </button>
      </div>

      <!-- Filter Kelas -->
      <div class="filter-container">
        <button
          v-for="k in kelasList"
          :key="k"
          class="btn btn-secondary m-1"
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
            <td @click="showModal(siswa)" class="text-primary cursor-pointer nama-siswa">
              {{ siswa.nama }}
            </td>
            <td>{{ siswa.jurusan }}</td>
            <td>{{ siswa.kelas }}</td>
            <td>{{ siswa.jumlahHadir }}</td>
            <td>{{ siswa.jumlahSakit }}</td>
            <td>{{ siswa.jumlahIzin }}</td>
            <td>{{ siswa.jumlahAlfa }}</td>
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
      jurusanList: ["RPL", "TKJ", "DPIB", "TBSM"],
      kelasList: ["A", "B", "C"],
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

.btn {
  margin-right: 5px;
}

.nama-siswa {
  font-weight: bold;
  font-size: 1.1rem;
  color: #007bff;
  cursor: pointer;
  transition: color 0.3s;
}

.nama-siswa:hover {
  color: #0056b3;
  text-decoration: none;
}
</style>
