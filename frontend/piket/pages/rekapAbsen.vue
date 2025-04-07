<template>
  <div class="container-absen">
    <div>
      <Navnew />
      <div class="container">

        <h1 style="margin-left: 380px;">Absen Siswa</h1>

        <!-- Filter -->
        <div class="filter-wrapper">
          <div class="filter-container">
            <label for="jurusan">Pilih Jurusan:</label>
            <select v-model="filterJurusan" id="jurusan" class="form-control">
              <option value="">Semua Jurusan</option>
              <option v-for="j in jurusanList" :key="j" :value="j">{{ j }}</option>
            </select>
          </div>

          <div class="filter-container">
            <label for="kelas">Pilih Kelas:</label>
            <select v-model="filterKelas" id="kelas" class="form-control">
              <option value="">Semua Kelas</option>
              <option v-for="k in kelasList" :key="k" :value="k">{{ k }}</option>
            </select>
          </div>

          <div class="filter-container">
            <label>&nbsp;</label>
            <button class="btn-reset" @click="resetFilter"><i class="fas fa-sync-alt"></i></button>
          </div>
        </div>

        <!-- Tabel -->
        <table class="table table-striped table-bordered">
          <thead>
            <tr>
              <th>No</th>
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
              <td>{{ index + 1 }}</td>
              <td>{{ siswa.nis }}</td>
              <td @click="showModal(siswa)" class="text-primary cursor-pointer nama-siswa">{{ siswa.nama }}</td>
              <td>{{ siswa.jurusan }}</td>
              <td>{{ siswa.kelas }}</td>
              <td>{{ siswa.jumlahHadir }}</td>
              <td>{{ siswa.jumlahSakit }}</td>
              <td>{{ siswa.jumlahIzin }}</td>
              <td>{{ siswa.jumlahAlfa }}</td>
            </tr>
          </tbody>
        </table>

        <!-- Aksi -->
        <div class="d-flex justify-center">
          <button class="button-download mr-1" @click="generatePDF">Download PDF</button>
          <button class="button-hapus" @click="hapusSemuaAbsensi" :disabled="loading">Hapus Semua Absensi</button>
        </div>

        <!-- Modal -->
        <RekapAbsensiModal
          :show="showModalFlag"
          :siswa="selectedSiswa"
          :absensi="absensiSiswa"
          @close="showModalFlag = false"
        />
      </div>
    </div>

<!-- Menu Grafik Gabungan Jurusan & Kelas -->
<div class="menu-gerafik">
  <div
    class="menu-item"
    v-for="jurusan in jurusanList"
    :key="jurusan"
  >
    <div v-for="kelas in kelasList" :key="jurusan + kelas" class="menu-link" style="cursor: default;">
      <strong>{{ jurusan }} - {{ kelas }}</strong>
      <ul style="list-style: none; padding-left: 0; margin-top: 8px; font-size: 14px;">
        <li>Hadir: {{ countByStatusGabungan(jurusan, kelas, 'hadir') }}</li>
        <li>Sakit: {{ countByStatusGabungan(jurusan, kelas, 'sakit') }}</li>
        <li>Izin: {{ countByStatusGabungan(jurusan, kelas, 'izin') }}</li>
        <li>Alfa: {{ countByStatusGabungan(jurusan, kelas, 'alfa') }}</li>
      </ul>
    </div>
  </div>
</div>


  </div>
</template>

<script>
import RekapAbsensiModal from "@/components/RekapAbsensiModal.vue";
import Navnew from "../components/Navnew.vue";
import axios from "axios";
import jsPDF from "jspdf";
import autoTable from "jspdf-autotable";

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
      filterJurusan: "",
      filterKelas: "",
      jurusanList: ["RPL", "MPLB", "PH"],
      kelasList: ["10A", "10B", "10C", "11A", "11B", "11C", "12A", "12B", "12C"],
      loading: false,
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

        await Promise.all(
          siswaData.map(async (siswa) => {
            try {
              const res = await axios.get(`http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(siswa.nama)}`);
              const absensi = res.data?.data || [];
              siswa.jumlahHadir = absensi.filter(a => a.status.toLowerCase() === "hadir").length;
              siswa.jumlahSakit = absensi.filter(a => a.status.toLowerCase() === "sakit").length;
              siswa.jumlahIzin = absensi.filter(a => a.status.toLowerCase() === "izin").length;
              siswa.jumlahAlfa = absensi.filter(a => a.status.toLowerCase() === "alfa").length;
            } catch {
              siswa.jumlahHadir = siswa.jumlahSakit = siswa.jumlahIzin = siswa.jumlahAlfa = 0;
            }
          })
        );

        this.rekapAbsensi = siswaData;
      } catch (error) {
        console.error("Gagal ambil data siswa:", error);
        this.rekapAbsensi = [];
      }
    },

    countByStatus(kelas, status) {
      return this.rekapAbsensi
        .filter(s => s.kelas === kelas)
        .reduce((total, s) => total + (s[`jumlah${this.capitalize(status)}`] || 0), 0);
    },

    countByStatusGabungan(jurusan, kelas, status) {
  return this.rekapAbsensi
    .filter(s => s.jurusan === jurusan && s.kelas === kelas)
    .reduce((total, s) => total + (s[`jumlah${this.capitalize(status)}`] || 0), 0);
}
,

    capitalize(word) {
      return word.charAt(0).toUpperCase() + word.slice(1);
    },



    async showModal(siswa) {
      this.selectedSiswa = siswa;
      this.showModalFlag = true;
      try {
        const response = await axios.get(`http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(siswa.nama)}`);
        this.absensiSiswa = response.data?.data || [];
      } catch (error) {
        console.error("Gagal ambil absensi:", error);
        this.absensiSiswa = [];
      }
    },

    async hapusSemuaAbsensi() {
      this.loading = true;
      try {
        await axios.delete("http://localhost:8080/api/absensi/all");
        this.$toast.success("Rekap absensi berhasil dihapus!", { position: "top-right", timeout: 3000 });
        await this.fetchAllSiswa();
      } catch (error) {
        this.$toast.error("Gagal menghapus absensi", { position: "top-right", timeout: 3000 });
      } finally {
        this.loading = false;
      }
    },

    generatePDF() {
      const doc = new jsPDF();
      doc.setFontSize(18);
      doc.text("Rekap Absensi Siswa", 14, 15);

      const headers = [["No", "NIS", "Nama", "Jurusan", "Kelas", "Hadir", "Sakit", "Izin", "Alfa"]];
      const data = this.filteredSiswa.map((siswa, index) => [
        index + 1,
        siswa.nis,
        siswa.nama,
        siswa.jurusan,
        siswa.kelas,
        siswa.jumlahHadir,
        siswa.jumlahSakit,
        siswa.jumlahIzin,
        siswa.jumlahAlfa,
      ]);

      autoTable(doc, {
        startY: 25,
        head: headers,
        body: data,
        styles: { fontSize: 10, cellPadding: 3 },
        headStyles: { fillColor: [41, 128, 185] },
      });

      doc.save("rekap_absensi.pdf");
    },

    resetFilter() {
      this.filterJurusan = "";
      this.filterKelas = "";
    },
  },
};
</script>

<style scoped>
.container-absen {
  display: flex;
  gap: 20px;
  justify-content: space-between;
  flex-wrap: nowrap;
  margin-top: 20px;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
}

.container {
  flex: 1;
  margin-top: 90px;
  background: #ffffff;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
  width: 1000px;
}

.menu-gerafik {
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 200px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  margin-top: 90px;
}

.menu-item {
  margin-bottom: 10px;
}

.menu-link {
  text-decoration: none;
  color: #333;
  padding: 10px;
  display: block;
  border-radius: 8px;
  transition: background 0.3s, color 0.3s;
}

.menu-link:hover {
  background-color: #007bff;
  color: #fff;
}

.filter-wrapper {
  display: flex;
  gap: 20px;
  margin-bottom: 20px;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
}

.filter-container {
  display: flex;
  flex-direction: column;
}

.form-control {
  padding: 8px;
  border-radius: 8px;
  border: 1px solid #ccc;
  min-width: 160px;
}

.btn-reset {
  background-color: transparent;
  color: #007bff;
  font-size: 24px;
  padding: 10px;
  border: 2px solid #007bff;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  width: 40px;
  height: 40px;
}

.btn-reset:hover {
  background-color: #5a6268;
}

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

.table tbody td {
  padding: 10px;
  text-align: center;
  border-bottom: 1px solid #ddd;
}

.table tbody tr:hover {
  background: #f2f2f2;
}

.nama-siswa {
  font-weight: bold;
  font-size: 1.1rem;
  color: black;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
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

.button-download {
  padding: 10px 15px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  margin-right: 10px;
}

.button-download:hover {
  background-color: #0056b3;
}

.button-hapus {
  padding: 10px 15px;
  background-color: red;
  color: white;
  border: none;
  border-radius: 10px;
  cursor: pointer;
}

.button-hapus:disabled {
  background-color: gray;
  cursor: not-allowed;
}

/* Responsive */
@media (max-width: 768px) {
  .container-absen {
    flex-direction: column;
    padding: 0 10px;
  }

  .menu-gerafik {
    width: 100%;
    flex-direction: row;
    justify-content: space-around;
    margin-top: 20px;
  }

  .menu-item {
    flex: 1;
    text-align: center;
  }

  .menu-link {
    padding: 10px 0;
  }
}
</style>
