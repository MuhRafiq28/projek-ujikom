<template>
  <div>
    <Navnew />
    <div class="container">
      <h2 class="title">Absensi Siswa</h2>

      <!-- Filter Dropdown for Jurusan and Kelas -->
      <div class="filter-container">
        <select
          v-model="selectedKelas"
          @change="fetchSiswa"
          class="form-select"
        >
          <option value="">Semua Kelas</option>
          <option v-for="kelas in kelasList" :key="kelas" :value="kelas">
            {{ kelas }}
          </option>
        </select>
        <select
          v-model="selectedJurusan"
          @change="fetchSiswa"
          class="form-select"
        >
          <option value="">Semua Jurusan</option>
          <option
            v-for="jurusan in jurusanList"
            :key="jurusan"
            :value="jurusan"
          >
            {{ jurusan }}
          </option>
        </select>
      </div>

      <table class="table">
        <thead>
          <tr>
            <th>No.</th>
            <th>Nama Siswa</th>
            <th>Kelas</th>
            <th>Jurusan</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(absensi, index) in filteredAbsensi" :key="index">
            <td>{{ index + 1 }}</td>
            <td>{{ absensi.nama }}</td>
            <td>{{ absensi.kelas }}</td>
            <td :class="getJurusanClass(absensi.jurusan)">
              {{ absensi.jurusan }}
            </td>
            <td>
              <select
                v-model="absensi.status"
                class="form-select"
                :class="getStatusClass(absensi.status)"
              >
                <option value="Hadir">Hadir</option>
                <option value="Sakit">Sakit</option>
                <option value="Izin">Izin</option>
                <option value="Alfa">Alfa</option>
              </select>
            </td>
          </tr>
        </tbody>
      </table>

      <button @click="submitAbsensi" class="btn-submit">Simpan Absensi</button>
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
      absensiList: [],
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
      jurusanList: ["RPL", "MPLB", "PH", "TO"],
      selectedKelas: "",
      selectedJurusan: "",
    };
  },
  computed: {
    filteredAbsensi() {
      return this.absensiList.filter((siswa) => {
        const matchesKelas =
          !this.selectedKelas || siswa.kelas === this.selectedKelas;
        const matchesJurusan =
          !this.selectedJurusan || siswa.jurusan === this.selectedJurusan;
        return matchesKelas && matchesJurusan;
      });
    },
  },
  methods: {
    async fetchSiswa() {
      try {
        const response = await axios.get("http://localhost:8080/api/siswa");
        this.absensiList = response.data.map((siswa) => ({
          siswa_id: siswa.id,
          nama: siswa.nama,
          kelas: siswa.kelas,
          jurusan: siswa.jurusan,
          status: "Hadir",
          tanggal: new Date().toISOString().split("T")[0],
        }));
      } catch (error) {
        console.error("Error fetching siswa:", error);
      }
    },
    getStatusClass(status) {
      if (status === "Sakit") return "status-sakit";
      if (status === "Izin") return "status-izin";
      if (status === "Alfa") return "status-alfa";
      return "";
    },
    getJurusanClass(jurusan) {
      switch (jurusan) {
        case "RPL":
          return "text-RPL";
        case "PH":
          return "text-PH";
        case "MPLB":
          return "text-MPLB";
        case "TO":
          return "text-TO";
        default:
          return "";
      }
    },
    async submitAbsensi() {
      try {
        console.log("Data yang dikirim:", this.absensiList);
        await axios.post("http://localhost:8080/api/absensi", this.absensiList);

        // Menampilkan notifikasi sukses
        this.$toast.success("Absensi berhasil disimpan!", {
          position: "top-right",
          timeout: 3000,
        });

        // Navigasi ke halaman home setelah sukses
       // this.$router.push("/rekapAbsen");
      } catch (error) {
        console.error("Error submitting absensi:", error);

        // Menampilkan notifikasi error
        this.$toast.error("Gagal menyimpan absensi", {
          position: "top-right",
          timeout: 3000,
        });
      }
    },
  },
  mounted() {
    this.fetchSiswa();
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
  margin-bottom: 20px;
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
  justify-content: center;
}

.form-select {
  padding: 10px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid #ddd;
  transition: all 0.3s ease;
}

.table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 30px;
}

.table th,
.table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
  font-size: 16px;
}

.table th {
  color: #d5d5d5;
  background: #96858f;
}

.table td {
  font-size: 14px;
}

/* Warna untuk teks jurusan */
.text-RPL {
  color: #007bff; /* Biru */
  font-weight: bold;
}

.text-PH {
  color: #28a745; /* Hijau */
  font-weight: bold;
}

.text-MPLB {
  color: #ff69b4; /* Pink */
  font-weight: bold;
}

.text-TO {
  color: #dc3545; /* Merah */
  font-weight: bold;
}

/* Warna berdasarkan status absensi */
.status-sakit {
  background-color: #f8d7da;
  color: #721c24;
}

.status-izin {
  background-color: #d1ecf1;
  color: #0c5460;
}

.status-alfa {
  background-color: #fff3cd;
  color: #856404;
}

.btn-submit {
  background-color: #28a745;
  color: white;
  padding: 12px 24px;
  font-size: 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  margin-top: 20px;
  transition: background-color 0.3s;
}

.btn-submit:hover {
  background-color: #218838;
}
</style>
