<template>
  <div>
    <Navnew />
    <div class="container">
      <h1>Rekap Absen Siswa</h1>

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
      <th>No</th> <!-- Mengubah ID menjadi No -->
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
    <tr v-for="(siswa, index) in filteredSiswa" :key="siswa.id">
      <td>{{ index + 1 }}</td> <!-- Nomor urut -->
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
import Navnew from "../components/Navnew.vue";
import axios from "axios";

export default {
  components: {
    RekapAbsensiModal,
    Navnew,
  },
  data() {
    return {
      rekapAbsensi: [],
      showModalFlag: false,
      selectedSiswa: null,
      absensiSiswa: [],
      filterJurusan: null,
      filterKelas: null,
      // Diubah sesuai permintaan
      jurusanList: ["RPL", "MPLB", "PH"],  // Diubah jurusannya
      kelasList: ["10A", "10B", "10C", "11A", "11B", "11C", "12A", "12B", "12C"],  // Filter kelas 10, 11, 12
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
/* Container styling */
.container {
  max-width: 1000px;
  margin: 0 auto;
  margin-top: 90px;
  background: #ffffff;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
}

/* Tombol filter */
.filter-container {
  margin-bottom: 15px;
}

.btn {
  padding: 8px 15px;
  font-size: 14px;
  font-weight: bold;
  border-radius: 8px;
  transition: all 0.3s ease-in-out;
}

.btn-primary {
  background: linear-gradient(135deg, #007bff, #0056b3);
  color: #fff;
  border: none;
}

.btn-secondary {
  background: linear-gradient(135deg, #28a745, #1c7a2e);
  color: #fff;
  border: none;
}

.btn:hover {
  transform: scale(1.05);
  opacity: 0.9;
}

/* Tabel styling */
.table {
  width: 100%;
  border-collapse: collapse;
  border-radius: 8px;
  overflow: hidden;
}

.table thead {
  background: linear-gradient(135deg, #9099A2, #96858F);
  color: white;
  font-weight: bold;
}

.table thead th {
  padding: 12px;
  text-align: center;
}

.table tbody tr {
  transition: background 0.3s ease-in-out;
}

.table tbody tr:hover {
  background: #f2f2f2;
}

/* Styling untuk nomor urut */
.table tbody td {
  padding: 10px;
  text-align: center;
  border-bottom: 1px solid #ddd;
}


/* Nama siswa dengan efek hover */
.nama-siswa {
  font-weight: bold;
  font-size: 1.1rem;
  color: black;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  text-decoration: none;
  position: relative;
}

.nama-siswa:hover {
  color: #606060;
}

.nama-siswa::after {
  content: "";
  display: block;
  width: 0;
  height: 2px;
  background: #a88f7f;
  transition: width 0.3s;
  position: absolute;
  bottom: -3px;
  left: 50%;
  transform: translateX(-50%);
}

.nama-siswa:hover::after {
  width: 100%;
}
</style>

