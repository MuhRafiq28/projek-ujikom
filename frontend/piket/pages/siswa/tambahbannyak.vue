<template>
  <div class="container">
    <h2>Tambah Data Siswa</h2>

    <!-- Input jumlah siswa -->
    <div class="input-group">
      <label>Jumlah Siswa:</label>
      <input type="number" v-model.number="jumlahSiswa" min="1" placeholder="Masukkan jumlah" />
      <button @click="generateForm">Generate Form</button>
    </div>

    <form @submit.prevent="submitForm" v-if="siswaList.length > 0">
      <div v-for="(siswa, index) in siswaList" :key="index" class="form-group">
        <input v-model="siswa.nis" placeholder="NIS" required />
        <input v-model="siswa.nama" placeholder="Nama" required />
        <input v-model="siswa.jurusan" placeholder="Jurusan" required />
        <input v-model="siswa.kelas" placeholder="Kelas" required />
        <button type="button" class="delete-btn" @click="removeSiswa(index)">Hapus</button>
      </div>
      <button type="submit" class="submit-btn">Simpan</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      jumlahSiswa: 1, // Default jumlah siswa
      siswaList: []
    };
  },
  methods: {
    generateForm() {
      if (this.jumlahSiswa > 0) {
        this.siswaList = Array.from({ length: this.jumlahSiswa }, () => ({
          nis: '',
          nama: '',
          jurusan: '',
          kelas: ''
        }));
      } else {
        alert('Masukkan jumlah siswa yang valid.');
      }
    },
    removeSiswa(index) {
      this.siswaList.splice(index, 1);
    },
    async submitForm() {
      try {
        const response = await axios.post('http://localhost:8080/api/siswa/batch', this.siswaList);
        alert(response.data.message);
      } catch (error) {
        alert('Terjadi kesalahan: ' + (error.response?.data?.error || error.message));
      }
    }
  }
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  color: #333;
}

.input-group {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
}

button {
  padding: 8px 12px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.3s;
}

button:hover {
  opacity: 0.8;
}

.form-group {
  display: flex;
  gap: 10px;
  align-items: center;
  margin-bottom: 10px;
}

.delete-btn {
  background: #e74c3c;
  color: white;
}

.submit-btn {
  width: 100%;
  background: #2ecc71;
  color: white;
  padding: 10px;
  font-size: 16px;
  margin-top: 10px;
}
</style>
