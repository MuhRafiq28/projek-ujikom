<template>
  <div v-if="isVisible" class="modal-overlay" @click.self="closeModal">
    <div class="modal-content">
      <h3 class="modal-title">Tambah Siswa</h3>

      <!-- Input jumlah siswa -->
      <div class="input-group">
        <label>Jumlah Siswa:</label>
        <div class="input-wrapper">
          <input type="number" v-model.number="jumlahSiswa" min="1" placeholder="Masukkan jumlah" />
          <button class="generate-btn" @click="generateForm">Set Jumlah</button>
          <button class="refresh-btn" @click="refreshForm"><i class="fas fa-sync-alt"></i></button>
        </div>
      </div>

      <!-- Input Jurusan dan Kelas hanya muncul jika form lebih dari 5 -->
      <div v-if="siswaList.length > 4" class="input-group">
        <label>Jurusan:</label>
        <input type="text" v-model="jurusan" placeholder="Masukkan jurusan" />

        <label>Kelas:</label>
        <input type="text" v-model="kelas" placeholder="Masukkan kelas" />

        <button class="apply-btn" @click="applyJurusanKelas">Terapkan ke Semua</button>
      </div>

      <form @submit.prevent="submitForm" v-if="siswaList.length > 0" class="siswa-form">
        <div class="form-grid" v-if="siswaList.length > 5">
          <div v-for="(siswa, index) in siswaList" :key="index" class="form-item">
            <input v-model="siswa.nis" placeholder="NIS" required />
            <input v-model="siswa.nama" placeholder="Nama" required />
            <input v-model="siswa.jurusan" placeholder="Jurusan" required />
            <input v-model="siswa.kelas" placeholder="Kelas" required />
            <button type="button" class="delete-btn" @click="removeSiswa(index)">Hapus</button>
          </div>
        </div>
        <div v-else>
          <div v-for="(siswa, index) in siswaList" :key="index" class="form-group">
            <input v-model="siswa.nis" placeholder="NIS" required />
            <input v-model="siswa.nama" placeholder="Nama" required />
            <input v-model="siswa.jurusan" placeholder="Jurusan" required />
            <input v-model="siswa.kelas" placeholder="Kelas" required />
            <button type="button" class="delete-btn" @click="removeSiswa(index)">Hapus</button>
          </div>
        </div>
        <button type="submit" class="submit-btn">Simpan</button>
      </form>

      <button class="close-btn" @click="closeModal">Tutup</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  props: {
    isVisible: Boolean
  },
  data() {
    return {
      jumlahSiswa: 1,
      siswaList: [],
      jurusan: '',
      kelas: ''
    };
  },
  methods: {
    generateForm() {
      if (this.jumlahSiswa > 0) {
        this.siswaList = Array.from({ length: this.jumlahSiswa }, () => ({
          nis: '',
          nama: '',
          jurusan: this.jurusan,
          kelas: this.kelas
        }));
      } else {
        alert('Masukkan jumlah siswa yang valid.');
      }
    },
    refreshForm() {
      this.siswaList = [];
      this.jumlahSiswa = 1;
      this.jurusan = '';
      this.kelas = '';
    },
    removeSiswa(index) {
      this.siswaList.splice(index, 1);
    },
    applyJurusanKelas() {
      this.siswaList.forEach(siswa => {
        siswa.jurusan = this.jurusan;
        siswa.kelas = this.kelas;
      });
    },
    async submitForm() {
      try {
        const response = await axios.post('http://localhost:8080/api/siswa/batch', this.siswaList);
        // Tampilkan notifikasi sukses
        this.$toast.success(response.data.message || "Data siswa berhasil ditambahkan!");
        this.$emit('refreshData');
        this.closeModal();
      } catch (error) {
        alert('Terjadi kesalahan: ' + (error.response?.data?.error || error.message));
      }
    },
    closeModal() {
      this.$emit('close');
    }
  }
};
</script>

<style scoped>
.refresh-btn {
  background-color: transparent;
  color: #007bff;
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.refresh-btn:hover {
  background-color: #ff9bee;
  transform: scale(1.1);
  transition: transform 0.2s ease-in-out, background-color 0.2s ease-in-out;
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
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  width: 900px;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  text-align: center;
}

.modal-title {
  font-size: 1.5rem;
  margin-bottom: 15px;
}

.input-group {
  margin-bottom: 15px;
  display: flex;
  justify-content: center; /* Pusatkan secara horizontal */
  align-items: center; /* Pusatkan secara vertikal */
}

.input-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
}

input {
  padding: 8px;
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.generate-btn {
  background-color: #007bff;
  color: white;
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.generate-btn:hover {
  background-color: #0056b3;
}

.apply-btn {
  background-color: #28a745;
  color: white;
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 8px;
}

.apply-btn:hover {
  background-color: #218838;
}

.form-group,
.form-item {
  margin-bottom: 10px;
  border: 3px solid #bc1ad9;
  padding: 10px;
  border-radius: 4px;
  background: #b300ff;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 10px;
}

.delete-btn {
  background-color: #ffffff;
  color: rgb(0, 0, 0);
  border: none;
  padding: 5px;
  cursor: pointer;
  border-radius: 4px;
  margin-top: 5px;
}

.delete-btn:hover {
  background-color: #ff0019;
  color: white;
}

.submit-btn,
.close-btn {
  background-color: #28a745;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 10px;
}

.submit-btn:hover {
  background-color: #218838;
}

.close-btn {
  background-color: #6c757d;
}

.close-btn:hover {
  background-color: #5a6268;
}
</style>
