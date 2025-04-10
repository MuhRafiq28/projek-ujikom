<template>
  <div class="container-absen">
    <div>
      <Navnew />
      <div class="container">
        <h1 >Absen Siswa</h1>

        <!-- Filter -->
        <div class="filter-wrapper">
          <div class="filter-container">
            <label for="jurusan">Pilih Jurusan:</label>
            <select v-model="filterJurusan" id="jurusan" class="form-control">
              <option value="">Semua Jurusan</option>
              <option v-for="j in jurusanList" :key="j" :value="j">
                {{ j }}
              </option>
            </select>
          </div>

          <div class="filter-container">
            <label for="kelas">Pilih Kelas:</label>
            <select v-model="filterKelas" id="kelas" class="form-control">
              <option value="">Semua Kelas</option>
              <option v-for="k in kelasList" :key="k" :value="k">
                {{ k }}
              </option>
            </select>
          </div>

          <div class="filter-container">
            <label for="kelas">Cari Nama Siswa:</label>

              <input
                type="text"
                v-model="filterNama"
                placeholder="Cari nama siswa..."
                class="input"
                style="border-radius: 5px;"
              />

          </div>

          <div class="filter-container">
            <label>&nbsp;</label>
            <button class="btn-reset" @click="resetFilter">
              <i class="fas fa-sync-alt"></i>
            </button>
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
              <td
                @click="showModal(siswa)"
                class="text-primary cursor-pointer nama-siswa"
              >
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

        <!-- Aksi -->
        <div class="d-flex justify-center">
          <button class="button-download mr-1" @click="generatePDF">
            Download PDF
          </button>
          <button
            class="button-hapus"
            @click="hapusSemuaAbsensi"
            :disabled="loading"
          >
            Hapus Semua Absensi
          </button>
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
      <!-- Sidebar statistik -->
      <div class="w-full lg:w-1/4 p-4 bg-gray-100 rounded-xl shadow-md">
        <i class="fas fa-chart-pie"></i>
        <span>Statistik Kehadiran</span>
        <div class="space-y-6 max-h-[600px] overflow-y-auto">
          <div v-for="(data, index) in dataStatistikGabungan" :key="index">
            <p class="font-semibold text-gray-700 mb-1">{{ data.label }}</p>
            <GrafikStatistik :dataStatistik="[data]" />
          </div>
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
import GrafikStatistik from "@/components/GrafikStatistik.vue";

export default {
  components: {
    RekapAbsensiModal,
    Navnew,
    GrafikStatistik,
  },
  data() {
    return {
      rekapAbsensi: [],
      showModalFlag: false,
      selectedSiswa: null,
      absensiSiswa: [],
      filterJurusan: "",
      filterKelas: "",
      filterNama: "",
      jurusanList: ["RPL", "MPLB", "PH"],
      kelasList: [
        "10A",
        "10B",
        "10C",
        "11A",
        "11B",
        "11C",
        "12A",
        "12B",
        "12C",
      ],
      loading: false,
    };
  },

  computed: {
    dataStatistikGabungan() {
      const statistik = [];
      this.filteredJurusanList.forEach((jurusan) => {
        this.filteredKelasList.forEach((kelas) => {
          statistik.push({
            jurusan,
            kelas,
            hadir: this.countByStatusGabungan(jurusan, kelas, "hadir"),
            sakit: this.countByStatusGabungan(jurusan, kelas, "sakit"),
            izin: this.countByStatusGabungan(jurusan, kelas, "izin"),
            alfa: this.countByStatusGabungan(jurusan, kelas, "alfa"),
          });
        });
      });
      return statistik;
    },
    filteredSiswa() {
      return this.rekapAbsensi.filter((siswa) => {
        const jurusanMatch =
          this.filterJurusan === "" || siswa.jurusan === this.filterJurusan;
        const kelasMatch =
          this.filterKelas === "" || siswa.kelas === this.filterKelas;
        const namaMatch =
          this.filterNama === "" ||
          siswa.nama.toLowerCase().includes(this.filterNama.toLowerCase());
        return jurusanMatch && kelasMatch && namaMatch;
      });
    },

    // Untuk grafik gabungan
    filteredJurusanList() {
      if (this.filterJurusan) return [this.filterJurusan];
      return this.jurusanList;
    },

    filteredKelasList() {
      if (this.filterKelas) return [this.filterKelas];
      return this.kelasList;
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
              const res = await axios.get(
                `http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(
                  siswa.nama
                )}`
              );
              const absensi = res.data?.data || [];
              siswa.jumlahHadir = absensi.filter(
                (a) => a.status.toLowerCase() === "hadir"
              ).length;
              siswa.jumlahSakit = absensi.filter(
                (a) => a.status.toLowerCase() === "sakit"
              ).length;
              siswa.jumlahIzin = absensi.filter(
                (a) => a.status.toLowerCase() === "izin"
              ).length;
              siswa.jumlahAlfa = absensi.filter(
                (a) => a.status.toLowerCase() === "alfa"
              ).length;
            } catch {
              siswa.jumlahHadir =
                siswa.jumlahSakit =
                siswa.jumlahIzin =
                siswa.jumlahAlfa =
                  0;
            }
          })
        );

        this.rekapAbsensi = siswaData;
      } catch (error) {
        console.error("Gagal ambil data siswa:", error);
        this.rekapAbsensi = [];
      }
    },

    countByStatusGabungan(jurusan, kelas, status) {
      return this.rekapAbsensi
        .filter((s) => s.jurusan === jurusan && s.kelas === kelas)
        .reduce(
          (total, s) => total + (s[`jumlah${this.capitalize(status)}`] || 0),
          0
        );
    },

    capitalize(word) {
      return word.charAt(0).toUpperCase() + word.slice(1);
    },

    async showModal(siswa) {
      this.selectedSiswa = siswa;
      this.showModalFlag = true;
      try {
        const response = await axios.get(
          `http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(
            siswa.nama
          )}`
        );
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
        this.$toast.success("Rekap absensi berhasil dihapus!", {
          position: "top-right",
          timeout: 3000,
        });
        await this.fetchAllSiswa();
      } catch (error) {
        this.$toast.error("Gagal menghapus absensi", {
          position: "top-right",
          timeout: 3000,
        });
      } finally {
        this.loading = false;
      }
    },

    generatePDF() {
      const doc = new jsPDF();
      doc.setFontSize(18);
      doc.text("Rekap Absensi Siswa", 14, 15);

      const headers = [
        [
          "No",
          "NIS",
          "Nama",
          "Jurusan",
          "Kelas",
          "Hadir",
          "Sakit",
          "Izin",
          "Alfa",
        ],
      ];
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

<!-- Styles tetap seperti sebelumnya -->

<style scoped>
/* === Global Container === */
.container-absen {
  display: flex;
  gap: 24px;
  padding: 20px;
  max-width: 1400px;
  margin: 0 auto;
  background: #f9fafb;
}

/* === Main Container (Tabel & Filter) === */
.container {
  flex: 1;
  margin-top: 90px;
  background: #ffffff;
  padding: 32px;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  min-width: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  margin-left: 20px;
}

h1 {
  font-size: 1.8rem;
  color: #2c3e50;
  margin-bottom: 28px;
  font-weight: 600;
  width: 100%;
  text-align: center;
}

/* === Filter Section (Improved) === */
.filter-wrapper {
  display: flex;
  gap: 16px;
  margin-bottom: 28px;
  flex-wrap: wrap;
  align-items: flex-end;
  width: 100%;
  max-width: 1000px;
  background: #f8fafc;
  padding: 16px;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
}

.filter-container {
  display: flex;
  flex-direction: column;
  min-width: 160px;
  flex-grow: 1;
}

.filter-container label {
  font-size: 0.875rem;
  color: #64748b;
  margin-bottom: 6px;
  font-weight: 500;
}

.form-control, .input {
  padding: 10px 12px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 0.875rem;
  background: white;
  transition: all 0.2s ease;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
  width: 100%;
}

.form-control:focus, .input:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

/* Perbaikan untuk select box */
select.form-control {
  height: 40px; /* Sesuaikan dengan tinggi input text */
  appearance: none; /* Hilangkan style default browser */
  -webkit-appearance: none;
  -moz-appearance: none;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3e%3cpolyline points='6 9 12 15 18 9'%3e%3c/polyline%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: right 10px center;
  background-size: 16px;
  padding-right: 30px; /* Beri ruang untuk ikon panah */
}

/* Pastikan semua input dan select memiliki tinggi yang sama */
.form-control, .input, select.form-control {
  min-height: 40px;
  box-sizing: border-box; /* Pastikan padding tidak mempengaruhi tinggi */
}

/* Hover state untuk select */
select.form-control:hover {
  border-color: #cbd5e1;
}

/* Focus state untuk select */
select.form-control:focus {
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

.btn-reset {
  background: #f1f5f9;
  color: #64748b;
  border: none;
  border-radius: 8px;
  padding: 10px 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40px;
  width: 40px;
}

.btn-reset:hover {
  background: #e2e8f0;
  color: #475569;
}

.btn-reset i {
  font-size: 14px;
}

/* === Table Section (Improved) === */
.table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 24px;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.table thead {
  background: #4f46e5;
  color: white;
  position: sticky;
  top: 0;
}

.table th, .table td {
  padding: 14px 12px;
  text-align: center;
  border-bottom: 1px solid #f1f5f9;
  word-wrap: break-word;
}

.table th {
  font-weight: 500;
  font-size: 0.875rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.table td {
  font-size: 0.875rem;
  color: #334155;
}

/* Kolom lebih proporsional */
.table th:nth-child(1), .table td:nth-child(1) { width: 7%; } /* No */
.table th:nth-child(2), .table td:nth-child(2) { width: 12%; } /* NIS */
.table th:nth-child(3), .table td:nth-child(3) { width: 20%; } /* Nama */
.table th:nth-child(4), .table td:nth-child(4) { width: 15%; } /* Jurusan */
.table th:nth-child(5), .table td:nth-child(5) { width: 10%; } /* Kelas */
.table th:nth-child(6), .table td:nth-child(6) { width: 10%; } /* Hadir */
.table th:nth-child(7), .table td:nth-child(7) { width: 10%; } /* Sakit */
.table th:nth-child(8), .table td:nth-child(8) { width: 10%; } /* Izin */
.table th:nth-child(9), .table td:nth-child(9) { width: 10%; } /* Alfa */

/* Hover effect */
.table tbody tr:hover {
  background: #f8fafc;
}

/* === Action Buttons === */
.button-container {
  display: flex;
  gap: 16px;
  margin-top: 24px;
  width: 100%;
  justify-content: center;
  max-width: 1000px;
}

.button-download {
  background: #4f46e5;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 10px 20px;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 500;
}

.button-download:hover {
  background: #4338ca;
}

.button-hapus {
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 10px 20px;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 500;
}

.button-hapus:hover {
  background: #dc2626;
}

.button-hapus:disabled {
  background: #fca5a5;
  cursor: not-allowed;
}

/* === Sidebar Grafik === */
.menu-gerafik {
  width: 320px;
  margin-top: 90px;
  background: #ffffff;
  padding: 24px;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  height: fit-content;
  margin-left: 30px;
}

.menu-gerafik i {
  margin-right: 8px;
  color: #4f46e5;
}

.menu-gerafik span {
  font-weight: 600;
  color: #2c3e50;
}

/* === Responsive Adjustments === */
@media (max-width: 1200px) {
  .container-absen {
    flex-direction: column;
    align-items: center;
  }

  .menu-gerafik {
    width: 100%;
    max-width: 1000px;
    margin-top: 32px;
  }
}

@media (max-width: 992px) {
  .container {
    padding: 28px 20px;
  }

  .table th:nth-child(3), .table td:nth-child(3) { width: 18%; } /* Nama */
  .table th:nth-child(4), .table td:nth-child(4) { width: 12%; } /* Jurusan */
}

@media (max-width: 768px) {
  .container {
    padding: 24px 16px;
  }

  .filter-wrapper {
    flex-direction: column;
    gap: 12px;
  }

  .filter-container {
    max-width: 100%;
  }

  .button-container {
    flex-direction: column;
    align-items: center;
  }

  /* Adjust table columns for mobile */
  .table th, .table td {
    padding: 12px 8px;
    font-size: 0.8rem;
  }

  .table th:nth-child(1), .table td:nth-child(1) { width: 8%; }
  .table th:nth-child(2), .table td:nth-child(2) { width: 12%; }
  .table th:nth-child(3), .table td:nth-child(3) { width: 22%; }
  .table th:nth-child(4), .table td:nth-child(4) { width: 14%; }
  .table th:nth-child(5), .table td:nth-child(5) { width: 12%; }
  .table th:nth-child(n+6), .table td:nth-child(n+6) { width: 8%; }
}
</style>
