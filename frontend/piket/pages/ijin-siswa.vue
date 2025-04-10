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
.container {
  max-width: 900px;
  margin: 90px auto 80px; /* Memberi jarak lebih dari navbar */
  padding: 20px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  gap: 20px;
}

.form-container {
  display: flex;
  align-items: center;
  width: 100%;
  gap: 30px;
}

.image-container img {
  width: 300px;
  border-radius: 10px;
}

.form-content {
  flex: 1;
  text-align: center;
}

/* Input dan select dengan efek hover */
input, select {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border-radius: 5px;
  border: 1px solid #ccc;
  font-size: 16px;
  transition: all 0.3s ease-in-out;
}

input:focus, select:focus {
  border-color: #8e44ad;
  box-shadow: 0 0 8px rgba(142, 68, 173, 0.5);
  transform: scale(1.05);
}

/* Animasi hover pada tombol */
button {
  width: 100%;
  padding: 12px;
  border-radius: 5px;
  border: none;
  font-size: 16px;
  color: white;
  background: linear-gradient(135deg, #3498db, #8e44ad);
  cursor: pointer;
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}

button:hover {
  transform: scale(1.1);
  box-shadow: 0 8px 15px rgba(142, 68, 173, 0.3);
}

button:active {
  transform: scale(1);
  box-shadow: 0 4px 8px rgba(142, 68, 173, 0.2);
}

/* Modal Styling */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background: white;
  padding: 15px 25px;
  border-radius: 8px;
  text-align: center;
  width: 300px;
  animation: fadeIn 0.2s ease-in-out;
}

.modal-content p {
  font-size: 16px;
  margin-bottom: 20px;
}

.modal-content button {
  background-color: #3498db;
  color: white;
  padding: 8px 16px;
  border-radius: 5px;
  border: none;
  cursor: pointer;
}

.modal-content button:hover {
  background-color: #8e44ad;
}

@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.9); }
  to { opacity: 1; transform: scale(1); }
}
</style>
