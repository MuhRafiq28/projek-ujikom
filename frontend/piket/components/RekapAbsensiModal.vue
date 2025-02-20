<template>
  <div v-if="show" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-content">
      <h3 class="mb-3">Rekap Absensi Siswa: {{ siswa.nama }}</h3>

      <div v-if="absensi.length === 0">
        <p class="text-danger">Tidak ada data absensi yang tersedia untuk siswa ini.</p>
      </div>

      <table v-else class="table table-bordered">
        <thead>
          <tr>
            <th>No</th>
            <th>Tanggal</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(rekap, index) in absensi" :key="rekap.id">
            <td>{{ index + 1 }}</td>
            <td>{{ formatDate(rekap.tanggal) }}</td>
            <td>{{ rekap.status }}</td>
          </tr>
        </tbody>
      </table>

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

        // Debugging: Cek isi response dari API
        console.log("Data absensi siswa:", response.data);

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
    formatDate(dateString) {
      if (!dateString) return "Tanggal tidak tersedia";
      const date = new Date(dateString);
      return date.toLocaleDateString("id-ID", {
        weekday: "short",
        year: "numeric",
        month: "numeric",
        day: "numeric"
      });
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
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 450px;
}

.text-danger {
  color: red;
}
</style>
