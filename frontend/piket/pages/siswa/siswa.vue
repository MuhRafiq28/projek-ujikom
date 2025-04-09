<template>
  <div>
    <Navnew />

    <div class="container" style="margin-top: 80px;">
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
        <div class="filter-group">
          <label>Cari Nama</label>
          <input type="text" v-model="searchNama" @input="fetchSiswa" placeholder="Cari berdasarkan nama..." />
        </div>
        <button @click="resetFilter" class="btn-refresh">
          <i class="fas fa-sync-alt"></i>
        </button>
      </div>

      <!-- Tambah Siswa Button -->
      <button @click="isModalVisible = true" class="btn-add">Tambah Siswa</button>
      <TambahBanyakSiswaModal :isVisible="isModalVisible" @close="isModalVisible = false" @refreshData="fetchSiswa" />

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
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(siswa, index) in siswaList" :key="siswa.id">
              <td>{{ index + 1 }}</td>
              <td>{{ siswa.nama }}</td>
              <td>{{ siswa.nis }}</td>
              <td>{{ siswa.jurusan }}</td>
              <td>{{ siswa.kelas }}</td>
              <td>
                <button @click="editSiswa(siswa)">Edit</button>
                <button @click="deleteSiswa(siswa.id)">Hapus</button>
              </td>
            </tr>
            <tr v-if="siswaList.length === 0">
              <td colspan="6">Tidak ada data siswa.</td>
            </tr>
          </tbody>
        </table>
        <!-- Rekap Jumlah Siswa per Kelas -->
<div class="rekap-jumlah">
  <h4>Jumlah Siswa per Kelas</h4>
  <ul>
    <li v-for="(jumlah, key) in jumlahPerKelas" :key="key">
      {{ key }}: {{ jumlah }}
    </li>
  </ul>
  <p><strong>Total keseluruhan siswa:</strong> {{ siswaList.length }}</p>
</div>

      </div>

      <!-- Modal Tambah/Edit -->
      <div v-if="isFormVisible" class="modal-overlay">
        <div class="modal-content">
          <h3>{{ formTitle }}</h3>
          <form @submit.prevent="submitForm">
            <input type="text" v-model="formData.nama" placeholder="Nama" required />
            <input type="text" v-model="formData.nis" placeholder="NIS" required />
            <select v-model="formData.jurusan">
              <option v-for="jurusan in jurusanList" :key="jurusan" :value="jurusan">
                {{ jurusan }}
              </option>
            </select>
            <select v-model="formData.kelas">
              <option v-for="kelas in kelasList" :key="kelas" :value="kelas">
                {{ kelas }}
              </option>
            </select>
            <button type="submit">Simpan</button>
            <button type="button" @click="closeForm">Batal</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from '@nuxtjs/composition-api'
import axios from 'axios'
import TambahBanyakSiswaModal from '@/components/TambahBanyakSiswaModal.vue'
import Navnew from '../../components/Navnew.vue'

export default {
  components: {
    TambahBanyakSiswaModal,
    Navnew
  },
  setup() {
    const siswaList = ref([])
    const selectedJurusan = ref('')
    const selectedKelas = ref('')
    const searchNama = ref('')
    const isModalVisible = ref(false)
    const isFormVisible = ref(false)
    const isEditing = ref(false)

    const formData = ref({
      id: null,
      nama: '',
      nis: '',
      jurusan: '',
      kelas: ''
    })

    const jurusanList = ref(['RPL', 'MPLB', 'PH', 'TO'])
    const kelasList = ref([
      '10A', '10B', '10C',
      '11A', '11B', '11C',
      '12A', '12B', '12C'
    ])

    const formTitle = computed(() => (isEditing.value ? 'Edit Siswa' : 'Tambah Siswa'))

    const fetchSiswa = async () => {
      try {
        let params = {}
        if (selectedJurusan.value) params.jurusan = selectedJurusan.value
        if (selectedKelas.value) params.kelas = selectedKelas.value
        if (searchNama.value) params.nama = searchNama.value

        const response = await axios.get('http://localhost:8080/api/siswa', { params })
        siswaList.value = response.data || []
      } catch (error) {
        console.error('Gagal mengambil data:', error)
      }
    }

    const addSiswa = () => {
      formData.value = { id: null, nama: '', nis: '', jurusan: '', kelas: '' }
      isFormVisible.value = true
      isEditing.value = false
    }

    const editSiswa = (siswa) => {
      formData.value = { ...siswa }
      isFormVisible.value = true
      isEditing.value = true
    }

    const deleteSiswa = async (id) => {
      if (!confirm('Apakah Anda yakin ingin menghapus siswa ini?')) return
      try {
        await axios.delete(`http://localhost:8080/api/siswa/${id}`)
        fetchSiswa()
      } catch (error) {
        console.error('Gagal menghapus data:', error)
      }
    }

    const jumlahPerKelas = computed(() => {
  const count = {}
  siswaList.value.forEach((siswa) => {
    const key = `${siswa.jurusan} ${siswa.kelas}`
    count[key] = (count[key] || 0) + 1
  })

  const ordered = {}
  urutanKelas.forEach((kelas) => {
    ordered[kelas] = count[kelas] || 0
  })
  return ordered
})



    const submitForm = async () => {
      try {
        if (isEditing.value) {
          const response = await axios.put(`http://localhost:8080/api/siswa/${formData.value.id}`, formData.value)
          siswaList.value = siswaList.value.map(siswa =>
            siswa.id === formData.value.id ? { ...siswa, ...formData.value } : siswa
          )
          alert(response.data.message || "Data siswa berhasil diperbarui!")
        } else {
          const response = await axios.post("http://localhost:8080/api/siswa", formData.value)
          siswaList.value.push(response.data)
          alert("Siswa berhasil ditambahkan!")
        }
        closeForm()
      } catch (error) {
        console.error("Gagal menyimpan data:", error)
        alert("Gagal menyimpan data!")
      }
    }

    const closeForm = () => {
      isFormVisible.value = false
    }

    const resetFilter = () => {
      selectedJurusan.value = ''
      selectedKelas.value = ''
      searchNama.value = ''
      fetchSiswa()
    }



    const urutanKelas = [
  'RPL 10A', 'RPL 10B', 'RPL 10C',
  'RPL 11A', 'RPL 11B', 'RPL 11C',
  'RPL 12A', 'RPL 12B', 'RPL 12C',
  'PH 10A', 'PH 10B', 'PH 10C',
  'PH 11A', 'PH 11B', 'PH 11C',
  'PH 12A', 'PH 12B', 'PH 12C',
  'MPLB 10A', 'MPLB 10B', 'MPLB 10C',
  'MPLB 11A', 'MPLB 11B', 'MPLB 11C',
  'MPLB 12A', 'MPLB 12B', 'MPLB 12C',
  'TO 10A', 'TO 10B', 'TO 10C',
  'TO 11A', 'TO 11B', 'TO 11C',
  'TO 12A', 'TO 12B', 'TO 12C'
]


    onMounted(fetchSiswa)

    return {
      siswaList,
      selectedJurusan,
      selectedKelas,
      searchNama,
      jurusanList,
      kelasList,
      isModalVisible,
      isFormVisible,
      isEditing,
      formData,
      formTitle,
      fetchSiswa,
      addSiswa,
      editSiswa,
      deleteSiswa,
      submitForm,
      closeForm,
      resetFilter,

  siswaList,
  jumlahPerKelas,


    }


  }
}
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
  z-index: 1000; /* Ensure modal is in front of other elements */
}


.form-container {
  margin-top: 20px;
}

.form-group {
  margin-bottom: 10px;
}

input, select {
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  width: 100%;
  font-size: 14px;
  box-sizing: border-box;
  transition: border-color 0.3s ease;
}

input:focus, select:focus {
  border-color: #007bff;
  outline: none;
}

button {
  background: #007bff;
  color: white;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background: #0056b3;
}

/* Filter Styling */
.filter-container {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  margin-bottom: 20px;
}

.filter-group {
  width: 48%;
}

.filter-group label {
  font-weight: bold;
  margin-bottom: 8px;
  display: block;
  font-size: 14px;
}

.filter-group select {
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  width: 100%;
  font-size: 14px;
  box-sizing: border-box;
  transition: border-color 0.3s ease;
}

.filter-group select:focus {
  border-color: #007bff;
  outline: none;
}

/* Tombol Reset Filter dengan Ikon */
.btn-refresh {
  background-color: transparent;
  color: #007bff;
  font-size: 24px;
  padding: 10px;
  border: 2px solid #007bff;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  width: 48px;
  height: 48px;
  margin-top: 25px;
}

.btn-refresh:hover {
  background-color: #007bff;
  color: white;
  border-color: #0056b3;
}

.btn-refresh i {
  margin: 0;
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
  width: 450px;
  position: relative;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: #dc3545;
  color: white;
  border: none;
  padding: 5px 10px;
  border-radius: 50%;
  cursor: pointer;
}

.close-btn:hover {
  background-color: #c82333;
}

/* Table Styling */
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.table-container {
  display: flex;

  gap:10px
}

table thead {
  background: linear-gradient(135deg, #9099A2, #96858F);
  color: white;
  font-weight: bold;
}

thead {
  background: #007bff;
  color: white;
}

th, td {
  padding: 12px;
  text-align: center;
  border: 1px solid #ddd;
}

tbody tr:hover {
  background: #f1f1f1;
}

button {
  padding: 5px 10px;
  margin: 0 5px;
  cursor: pointer;
}

button:hover {
  opacity: 0.8;
}

.btn-add {
  background-color: #28a745;
  color: white;
  font-size: 16px;
  font-weight: bold;
  padding: 10px 15px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  margin-bottom: 20px;
}

.btn-add:hover {
  background-color: #218838;
}

/* Search input style tambahan */
.filter-group input[type="text"] {
  padding: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  width: 100%;
  font-size: 14px;
  box-sizing: border-box;
  transition: border-color 0.3s ease;
}

.filter-group input[type="text"]:focus {
  border-color: #007bff;
  outline: none;
}

.rekap-jumlah {
  margin-top: 20px;
  background: #f0f0f0;
  padding: 16px;
  border-radius: 8px;
  max-width: 300px;
}
.rekap-jumlah h4 {
  margin-bottom: 10px;
  font-size: 16px;
}
.rekap-jumlah li {
  margin-bottom: 6px;
  font-size: 14px;
}

/* (Style lainnya tidak berubah, tetap seperti sebelumnya) */
</style>
