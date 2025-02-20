<template>
  <div>
    <Navbar />
    <div class="container">
      <h2 class="title">Absensi Siswa</h2>

      <table class="table">
        <thead>
          <tr>
            <th>Nama Siswa</th>
            <th>Kelas</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(absensi, index) in absensiList" :key="index">
            <td>{{ absensi.nama }}</td>
            <td>{{ absensi.kelas }}</td>
            <td>
              <select v-model="absensi.status" class="form-select">
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

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import Navbar from '../components/Navbar.vue';

const absensiList = ref([]);

// Ambil daftar siswa dan set default status Hadir
const fetchSiswa = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/siswa');
    absensiList.value = response.data.map((siswa) => ({
      siswa_id: siswa.id,
      nama: siswa.nama,
      kelas: siswa.kelas,
      status: 'Hadir', // Default status diubah ke Hadir
      tanggal: new Date().toISOString().split('T')[0], // Set tanggal otomatis
    }));
  } catch (error) {
    console.error('Error fetching siswa:', error);
  }
};

// Kirim data absensi ke backend
const submitAbsensi = async () => {
  try {
    console.log("Data yang dikirim:", absensiList.value);  // Debugging data yang dikirim
    await axios.post('http://localhost:8080/api/absensi', absensiList.value);  // Kirim seluruh data absensi
    alert('Absensi berhasil disimpan!');
  } catch (error) {
    console.error('Error submitting absensi:', error);
    alert('Gagal menyimpan absensi');
  }
};

onMounted(fetchSiswa);
</script>

<style scoped>
/* Container */
.container {
  max-width: 800px;
  margin-top: 50px;
  padding: 20px;
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
}

/* Judul */
.title {
  text-align: center;
  color: #333;
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 20px;
}

/* Tabel */
.table {
  width: 100%;
  border-collapse: collapse;
}

.table th, .table td {
  padding: 10px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

/* Tombol Submit */
.btn-submit {
  background-color: #007bff;
  color: #fff;
  font-size: 16px;
  font-weight: bold;
  padding: 12px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s ease;
  margin-top: 20px;
  display: block;
  width: 100%;
}

.btn-submit:hover {
  background-color: #0056b3;
}
</style>
