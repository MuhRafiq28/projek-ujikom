<template>
  <div class="container">
    <h2 class="title">Data Siswa</h2>

    <!-- Filter -->
    <div class="filter-container">
      <div class="filter-group">
        <label>Jurusan</label>
        <select v-model="selectedJurusan" @change="fetchSiswa">
          <option value="">Semua</option>
          <option v-for="jurusan in jurusanList" :key="jurusan" :value="jurusan">
            {{ jurusan }}
          </option>
        </select>
      </div>
      <div class="filter-group">
        <label>Kelas</label>
        <select v-model="selectedKelas" @change="fetchSiswa">
          <option value="">Semua</option>
          <option v-for="kelas in kelasList" :key="kelas" :value="kelas">
            {{ kelas }}
          </option>
        </select>
      </div>
    </div>

    <!-- Tabel -->
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>#</th>
            <th>Nama</th>
            <th>NIS</th>
            <th>Jurusan</th>
            <th>Kelas</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(siswa, index) in siswaList" :key="siswa.id">
            <td>{{ index + 1 }}</td>
            <td>{{ siswa.nama }}</td>
            <td>{{ siswa.nis }}</td>
            <td>{{ siswa.jurusan }}</td>
            <td>{{ siswa.kelas }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const siswaList = ref([])
const selectedJurusan = ref('')
const selectedKelas = ref('')
const jurusanList = ['RPL', 'TKJ', 'MM', 'DKV']
const kelasList = ['A', 'B', 'C']

const fetchSiswa = async () => {
  try {
    let url = 'http://localhost:8080/api/siswa' // Pakai /api/siswa
    let params = {}

    if (selectedJurusan.value) params.jurusan = selectedJurusan.value
    if (selectedKelas.value) params.kelas = selectedKelas.value

    const response = await axios.get(url, { params })
    console.log('Data siswa:', response.data) // Debugging
    siswaList.value = response.data
  } catch (error) {
    console.error('Gagal mengambil data:', error.response ? error.response.data : error.message)
  }
}

onMounted(fetchSiswa)
</script>

<style scoped>
/* Container */
.container {
  max-width: 900px;
  margin: auto;
  padding: 20px;
}

/* Title */
.title {
  text-align: center;
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

/* Filter */
.filter-container {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-bottom: 20px;
}

.filter-group {
  display: flex;
  flex-direction: column;
}

.filter-group label {
  font-weight: bold;
  margin-bottom: 5px;
}

.filter-group select {
  padding: 8px;
  border-radius: 5px;
  border: 1px solid #ccc;
  cursor: pointer;
}

.filter-group select:hover {
  border-color: #007bff;
}

/* Table */
.table-container {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  background: #fff;
  border-radius: 5px;
  overflow: hidden;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

thead {
  background: #007bff;
  color: white;
}

th, td {
  padding: 12px;
  text-align: center;
  border-bottom: 1px solid #ddd;
}

tbody tr:hover {
  background: #f1f1f1;
}
</style>
