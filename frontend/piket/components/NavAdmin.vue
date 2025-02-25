<template>
  <div :class="userRole === 'admin' ? 'admin-layout' : 'user-layout'">
    <NavAdmin v-if="userRole === 'admin'" />
    <div class="main-content">
      <Navbar v-if="userRole !== 'admin'" />

      <div class="container">
        <h2 class="title">Absensi Siswa</h2>

        <!-- Filter Dropdown for Jurusan dan Kelas -->
        <div class="filter-container">
          <select v-model="selectedKelas" @change="fetchSiswa" class="form-select">
            <option value="">Semua Kelas</option>
            <option v-for="kelas in kelasList" :key="kelas" :value="kelas">{{ kelas }}</option>
          </select>
          <select v-model="selectedJurusan" @change="fetchSiswa" class="form-select">
            <option value="">Semua Jurusan</option>
            <option v-for="jurusan in jurusanList" :key="jurusan" :value="jurusan">{{ jurusan }}</option>
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
              <td>{{ absensi.jurusan }}</td>
              <td>
                <select v-model="absensi.status" class="form-select" :class="getStatusClass(absensi.status)">
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
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';
import Navbar from '../components/Navbar.vue';
import NavAdmin from '../components/NavAdmin.vue';

// State
const absensiList = ref([]);
const kelasList = ['10A', '10B', '10C', '11A', '11B', '11C', '12A', '12B', '12C'];
const jurusanList = ['RPL', 'MPLB', 'PH'];
const selectedKelas = ref('');
const selectedJurusan = ref('');
const userRole = ref('');

// Fetch data siswa
const fetchSiswa = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/siswa');
    absensiList.value = response.data.map((siswa) => ({
      siswa_id: siswa.id,
      nama: siswa.nama,
      kelas: siswa.kelas,
      jurusan: siswa.jurusan,
      status: 'Hadir',
      tanggal: new Date().toISOString().split('T')[0],
    }));
  } catch (error) {
    console.error('Error fetching siswa:', error);
  }
};

// Filter absensi berdasarkan kelas dan jurusan
const filteredAbsensi = computed(() => {
  return absensiList.value.filter(siswa => {
    const matchesKelas = !selectedKelas.value || siswa.kelas === selectedKelas.value;
    const matchesJurusan = !selectedJurusan.value || siswa.jurusan === selectedJurusan.value;
    return matchesKelas && matchesJurusan;
  });
});

// Warna status
const getStatusClass = (status) => {
  if (status === 'Sakit') return 'status-sakit';
  if (status === 'Izin') return 'status-izin';
  if (status === 'Alfa') return 'status-alfa';
  return '';
};

// Submit absensi
const submitAbsensi = async () => {
  try {
    console.log("Data yang dikirim:", absensiList.value);
    await axios.post('http://localhost:8080/api/absensi', absensiList.value);
    alert('Absensi berhasil disimpan!');
  } catch (error) {
    console.error('Error submitting absensi:', error);
    alert('Gagal menyimpan absensi');
  }
};

// Set userRole dari localStorage
onMounted(() => {
  const role = localStorage.getItem('userRole');
  userRole.value = role ? role : '';
  fetchSiswa();
});
</script>

<style scoped>
/* Layout untuk Admin */
.admin-layout {
  display: flex;
}

.admin-layout .main-content {
  margin-left: 250px; /* Geser konten ke kanan */
  width: calc(100% - 250px);
  padding: 20px;
}

/* Layout untuk User */
.user-layout .main-content {
  margin-top: 80px; /* Pastikan tidak tertutup navbar */
  padding: 20px;
}

/* Sidebar Styling */
.sidebar {
  width: 250px;
  height: 100vh;
  background: #343a40;
  color: white;
  position: fixed;
  padding: 20px;
}

/* Container */
.container {
  max-width: 1000px;
  margin: 0 auto;
}

/* Filter */
.filter-container {
  margin-bottom: 20px;
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
  justify-content: center;
}

/* Form Select */
.form-select {
  padding: 10px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid #ddd;
  transition: all 0.3s ease;
}

/* Table */
.table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 30px;
}

.table th, .table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
  font-size: 16px;
}

.table th {
  background-color: #f4f4f4;
  color: #555;
}

.table td {
  font-size: 14px;
}

/* Status Colors */
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

/* Submit Button */
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
