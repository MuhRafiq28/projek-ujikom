<template>
  <div v-if="show" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <h3 class="mb-3">Rekap Absensi Siswa: {{ siswa.nama }}</h3>

      <div v-if="absensi.length === 0">
        <p class="text-danger">Tidak ada data absensi yang tersedia untuk siswa ini.</p>
      </div>

      <div v-else class="absensi-container">
        <div v-for="(rekap, index) in absensi" :key="rekap.id" class="absensi-card">
          <p class="nomor">{{ index + 1 }}</p>
          <p class="tanggal">{{ formatDate(rekap.tanggal) }}</p>
          <p :class="['status', getStatusClass(rekap.status)]">{{ rekap.status }}</p>
          <button class="btn btn-danger btn-sm" @click="deleteAbsensi(rekap.id)">Hapus</button>
        </div>
      </div>

      <button class="btn btn-danger mt-3" v-if="absensi.length" @click="deleteAllAbsensi">Hapus Semua Absensi</button>
      <button class="btn btn-secondary mt-3" @click="$emit('close')">Tutup</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  props: {
    show: Boolean,
    siswa: Object,
  },
  data() {
    return {
      absensi: [],
    };
  },
  watch: {
    show(newVal) {
      if (newVal) {
        this.fetchAbsensi();
      }
    },
  },
  methods: {
    async fetchAbsensi() {
      if (!this.siswa || !this.siswa.nama) return;
      try {
        const response = await axios.get(`http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(this.siswa.nama)}`);
        if (response.data.status === "success" && Array.isArray(response.data.data)) {
          this.absensi = response.data.data;
        } else {
          this.absensi = [];
        }
      } catch (error) {
        console.error("Gagal mengambil data absensi:", error);
        this.absensi = [];
      }
    },
    async deleteAbsensi(id) {
      if (!confirm("Yakin ingin menghapus absensi ini?")) return;
      try {
        await axios.delete(`http://localhost:8080/api/rekap-absensi/${id}`);
        this.absensi = this.absensi.filter(item => item.id !== id);
      } catch (error) {
        console.error("Gagal menghapus absensi:", error);
      }
    },
    async deleteAllAbsensi() {
      if (!confirm("Yakin ingin menghapus semua absensi siswa ini?")) return;
      try {
        await axios.delete(`http://localhost:8080/api/rekap-absensi/nama/${encodeURIComponent(this.siswa.nama)}`);
        this.absensi = [];
      } catch (error) {
        console.error("Gagal menghapus semua absensi:", error);
      }
    },
    formatDate(dateString) {
      if (!dateString) return "Tanggal tidak tersedia";
      const date = new Date(dateString);
      return date.toLocaleDateString("id-ID", {
        weekday: "short",
        year: "numeric",
        month: "numeric",
        day: "numeric"
      });
    },
    getStatusClass(status) {
      switch (status.toLowerCase()) {
        case 'sakit': return 'sakit';
        case 'izin': return 'izin';
        case 'alfa': return 'alfa';
        case 'hadir': return 'hadir';
        default: return '';
      }
    }
  }
};
</script>


<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1050; /* Pastikan modal berada di atas navbar */
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 700px; /* Modal lebih lebar */
}

.text-danger {
  color: red;
}

/* Tampilan absensi dalam bentuk grid */
.absensi-container {
  display: grid;
  grid-template-columns: repeat(5, 1fr); /* 5 kolom */
  gap: 10px;
  max-height: 400px; /* Sesuaikan dengan ukuran yang diinginkan */
  overflow-y: auto; /* Agar bisa scroll ke bawah jika data lebih banyak */
}

.absensi-card {
  background: #f8f9fa;
  padding: 10px;
  border-radius: 6px;
  text-align: center;
  min-height: 100px; /* Tentukan tinggi minimum untuk setiap card */
}

.nomor {
  font-weight: bold;
}

.tanggal {
  font-size: 0.9em;
  color: #555;
}

.status {
  font-size: 1em;
  font-weight: bold;
  padding: 5px;
  border-radius: 5px;
  color: white; /* Teks warna putih untuk kontras */
}

/* Warna berdasarkan status */
.sakit {
  background-color: blue;
}

.izin {
  background-color: yellow;
  color: black; /* Warna teks hitam untuk kontras dengan latar belakang kuning */
}

.alfa {
  background-color: red;
}

.hadir {
  background-color: green;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 700px; /* Modal lebih lebar */
}

.text-danger {
  color: red;
}

/* Tampilan absensi dalam bentuk grid */
.absensi-container {
  display: grid;
  grid-template-columns: repeat(5, 1fr); /* 5 kolom */
  gap: 10px;
  max-height: 400px; /* Sesuaikan dengan ukuran yang diinginkan */
  overflow-y: auto; /* Agar bisa scroll ke bawah jika data lebih banyak */
}

.absensi-card {
  background: #f8f9fa;
  padding: 10px;
  border-radius: 6px;
  text-align: center;
  min-height: 100px; /* Tentukan tinggi minimum untuk setiap card */
}

.nomor {
  font-weight: bold;
}

.tanggal {
  font-size: 0.9em;
  color: #555;
}

.status {
  font-size: 1em;
  font-weight: bold;
  padding: 5px;
  border-radius: 5px;
  color: white; /* Teks warna putih untuk kontras */
}

/* Warna berdasarkan status */
.sakit {
  background-color: blue;
}

.izin {
  background-color: yellow;
  color: black; /* Warna teks hitam untuk kontras dengan latar belakang kuning */
}

.alfa {
  background-color: red;
}

.hadir {
  background-color: green;
}
</style>
