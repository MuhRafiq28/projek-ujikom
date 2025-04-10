<template>
  <div>
    <Navnew />
    <div class="container">
      <div class="form-container">
        <div class="image-container">
          <img src="/images/inspirasi-logo.png" alt="Izin Illustration" />
        </div>
        <div class="form-content">
          <h1>Buat Izin</h1>
          <form @submit.prevent="tambahIzin">
            <input v-model="izinBaru.nama" placeholder="Nama" required />
            <input v-model="izinBaru.alasan" placeholder="Alasan" required />

            <select v-model="izinBaru.jurusan" required>
              <option disabled value="">Pilih Jurusan</option>
              <option value="RPL">RPL</option>
              <option value="PH">PH</option>
              <option value="MPLB">MPLB</option>
              <option value="TO">TO</option>
            </select>

            <select v-model="izinBaru.kelas" required>
              <option disabled value="">Pilih Kelas</option>
              <option value="10A">10A</option>
              <option value="10B">10B</option>
              <option value="10C">10C</option>
              <option value="11A">11A</option>
              <option value="11B">11B</option>
              <option value="11C">11C</option>
              <option value="12A">12A</option>
              <option value="12B">12B</option>
              <option value="12C">12C</option>
            </select>

            <select v-model="izinBaru.status" required>
              <option value="Keluar">Keluar</option>
              <option value="Masuk">Masuk</option>
              <option value="Terlambat">Terlambat</option>
              <option value="Pulang Cepat">Pulang Lebih Cepat</option>
            </select>

            <button type="submit">Tambah Izin</button>
          </form>
        </div>
      </div>
    </div>

    <!-- Modal Konfirmasi -->
    <div v-if="showModal" class="modal-overlay">
      <div class="modal-content">
        <p>Apakah kamu yakin ingin menyimpan data izin ini?</p>
        <div style="display: flex; justify-content: center; gap: 10px">
          <button @click="showModal = false">Batal</button>
          <button @click="konfirmasiTambahIzin">Ya, Simpan</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Navnew from "../components/Navnew.vue";

export default {
  components: { Navnew },
  data() {
    return {
      izinBaru: {
        nama: "",
        alasan: "",
        jurusan: "",
        kelas: "",
        status: "Keluar",
        waktu_keluar: ""  // Menambahkan field waktu_keluar
      },
      showModal: false,
      modalMessage: "",
    };
  },

  methods: {
    tambahIzin() {
      // Menambahkan waktu_keluar saat status 'Pulang Lebih Cepat'
      if (this.izinBaru.status === "Pulang Lebih Cepat") {
        this.izinBaru.waktu_keluar = new Date().toISOString(); // Menyimpan waktu saat izin dibuat
      }

      // Menampilkan modal konfirmasi
      this.showModal = true;
    },

    async konfirmasiTambahIzin() {
      this.showModal = false;

      // Menyiapkan data yang akan dikirim sesuai format
      const izinData = {
        nama: this.izinBaru.nama,
        alasan: this.izinBaru.alasan,
        jurusan: this.izinBaru.jurusan,
        kelas: this.izinBaru.kelas,
        status: this.izinBaru.status === "Pulang Cepat" ? "Pulang Lebih Cepat" : this.izinBaru.status, // Mengubah status menjadi 'Pulang Lebih Cepat' jika 'Pulang Cepat'
        // hanya menambahkan waktu_keluar jika statusnya "Pulang Lebih Cepat"
        ...(this.izinBaru.status === "Pulang Lebih Cepat" && { waktu_keluar: this.izinBaru.waktu_keluar })
      };

      try {
        // Kirim data izin ke backend
        const response = await axios.post("http://localhost:8080/api/izin", izinData);

        if (response.data.message === "Izin berhasil ditambahkan") {
          this.$toast.success("Izin berhasil disimpan!", {
            position: "top-right",
            timeout: 3000,
          });

          // Reset form setelah berhasil
          this.izinBaru = {
            nama: "",
            alasan: "",
            jurusan: "",
            kelas: "",
            status: "Keluar",
            waktu_keluar: ""
          };
        }
      } catch (error) {
        this.$toast.error("Gagal menyimpan izin", {
          position: "top-right",
          timeout: 3000,
        });
      }
    }
  }
};
</script>

<style scoped>
/* === Global Styles === */
.container {
  max-width: 800px;
  margin: 90px auto 40px;
  padding: 0;
}

/* === Form Layout === */
.form-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40px;
  align-items: center;
}

.image-container img {
  width: 100%;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.form-content {
  padding: 30px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
}

h1 {
  font-size: 1.8rem;
  color: #2c3e50;
  margin-bottom: 24px;
  font-weight: 600;
  text-align: center;
}

/* === Form Elements === */
form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

input, select {
  width: 100%;
  padding: 12px 16px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 0.95rem;
  background: #f8fafc;
  transition: all 0.2s ease;
}

input:focus, select:focus {
  outline: none;
  border-color: #6366f1;
  background: white;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.1);
}

/* Custom Select Styling */
select {
  appearance: none;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%2364758b' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3e%3cpolyline points='6 9 12 15 18 9'%3e%3c/polyline%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: right 12px center;
  background-size: 16px;
  padding-right: 36px;
}

/* Button Styling */
button {
  padding: 14px;
  background: #6366f1;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  margin-top: 8px;
}

button:hover {
  background: #4f46e5;
  transform: translateY(-1px);
}

button:active {
  transform: translateY(0);
}

/* Modal Styling */
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
  padding: 24px;
  border-radius: 12px;
  width: 90%;
  max-width: 400px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  animation: modalFadeIn 0.2s ease-out;
}

.modal-content p {
  margin-bottom: 24px;
  color: #334155;
  text-align: center;
}

.modal-content div {
  display: flex;
  gap: 12px;
  justify-content: center;
}

.modal-content button {
  flex: 1;
  margin: 0;
  padding: 10px;
}

.modal-content button:first-child {
  background: #f1f5f9;
  color: #64748b;
}

.modal-content button:first-child:hover {
  background: #e2e8f0;
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
@media (max-width: 768px) {
  .form-container {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .image-container {
    display: none;
  }

  .form-content {
    padding: 24px;
  }

  h1 {
    font-size: 1.5rem;
  }
}
</style>
