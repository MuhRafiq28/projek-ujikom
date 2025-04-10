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
      <button @click="isModalVisible = true" class="btn-add-bulk">Tambah Banyak Siswa</button>
      <TambahBanyakSiswaModal :isVisible="isModalVisible" @close="isModalVisible = false" @refreshData="fetchSiswa" />

      <!-- Tabel -->
      <div class="table-container">
        <div class="table-wrapper">
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
                  <button @click="editSiswa(siswa)" class="btn-edit">Edit</button>
                  <button @click="deleteSiswa(siswa.id)" class="btn-delete">Hapus</button>
                </td>
              </tr>
              <tr v-if="siswaList.length === 0">
                <td colspan="6">Tidak ada data siswa.</td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Rekap Jumlah Siswa per Kelas -->
        <div class="rekap-jumlah">
          <div class="rekap-header">
            <h4><i class="fas fa-chart-pie"></i> Rekap Siswa</h4>
            <div class="total-siswa">
              <span>Total</span>
              <span class="total-number">{{ semuaSiswa.length }}</span>
            </div>
          </div>
          <div class="rekap-scroll-container">
            <div class="rekap-grid">
              <div v-for="(jumlah, key) in jumlahPerKelas" :key="key" class="rekap-item">
                <span class="kelas-label">{{ key.split(' ')[0] }}</span>
                <span class="kelas-romawi">{{ key.split(' ')[1] }}</span>
                <span class="jumlah-siswa">{{ jumlah }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Modal Tambah/Edit -->
      <div v-if="isFormVisible" class="modal-overlay">
        <div class="modal-content">
          <h3>{{ formTitle }}</h3>
          <form @submit.prevent="submitForm">
            <div class="form-group">
              <label>Nama</label>
              <input type="text" v-model="formData.nama" placeholder="Nama" required />
            </div>
            <div class="form-group">
              <label>NIS</label>
              <input type="text" v-model="formData.nis" placeholder="NIS" required />
            </div>
            <div class="form-group">
              <label>Jurusan</label>
              <select v-model="formData.jurusan" required>
                <option value="">Pilih Jurusan</option>
                <option v-for="jurusan in jurusanList" :key="jurusan" :value="jurusan">
                  {{ jurusan }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label>Kelas</label>
              <select v-model="formData.kelas" required>
                <option value="">Pilih Kelas</option>
                <option v-for="kelas in kelasList" :key="kelas" :value="kelas">
                  {{ kelas }}
                </option>
              </select>
            </div>
            <div class="form-actions">
              <button type="submit" class="btn-submit">Simpan</button>
              <button type="button" @click="closeForm" class="btn-cancel">Batal</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from '@nuxtjs/composition-api'
import axios from 'axios'
import Swal from 'sweetalert2'
import TambahBanyakSiswaModal from '@/components/TambahBanyakSiswaModal.vue'
import Navnew from '../../components/Navnew.vue'

export default {
  components: {
    TambahBanyakSiswaModal,
    Navnew
  },
  setup() {
    const siswaList = ref([])
    const semuaSiswa = ref([])
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
    const kelasList = ref(['10A', '10B', '10C', '11A', '11B', '11C', '12A', '12B', '12C'])

    const urutanKelas = [
      'RPL 10A', 'RPL 10B', 'RPL 10C', 'RPL 11A', 'RPL 11B', 'RPL 11C', 'RPL 12A', 'RPL 12B', 'RPL 12C',
      'PH 10A', 'PH 10B', 'PH 10C', 'PH 11A', 'PH 11B', 'PH 11C', 'PH 12A', 'PH 12B', 'PH 12C',
      'MPLB 10A', 'MPLB 10B', 'MPLB 10C', 'MPLB 11A', 'MPLB 11B', 'MPLB 11C', 'MPLB 12A', 'MPLB 12B', 'MPLB 12C',
      'TO 10A', 'TO 10B', 'TO 10C', 'TO 11A', 'TO 11B', 'TO 11C', 'TO 12A', 'TO 12B', 'TO 12C'
    ]

    const formTitle = computed(() => isEditing.value ? 'Edit Siswa' : 'Tambah Siswa')

    const hasChanges = (original, edited) => {
      return Object.keys(edited).some(key => 
        key !== 'id' && original[key] !== edited[key]
      )
    }

    const fetchSiswa = async () => {
      try {
        // Ambil semua data siswa tanpa filter untuk rekap
        const resAll = await axios.get('http://localhost:8080/api/siswa')
        semuaSiswa.value = resAll.data || []

        // Ambil data dengan filter untuk tabel
        let params = {}
        if (selectedJurusan.value) params.jurusan = selectedJurusan.value
        if (selectedKelas.value) params.kelas = selectedKelas.value
        if (searchNama.value) params.nama = searchNama.value

        const response = await axios.get('http://localhost:8080/api/siswa', { params })
        siswaList.value = response.data || []
      } catch (error) {
        console.error('Gagal mengambil data:', error)
        Swal.fire({
          icon: 'error',
          title: 'Gagal',
          text: 'Gagal memuat data siswa',
          confirmButtonColor: '#d33',
        })
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
      const siswa = siswaList.value.find(s => s.id === id)
      const { isConfirmed } = await Swal.fire({
        title: 'Konfirmasi Hapus',
        text: `Apakah Anda yakin ingin menghapus siswa ${siswa.nama}?`,
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: 'Ya, Hapus!',
        cancelButtonText: 'Batal'
      })

      if (!isConfirmed) return

      try {
        await axios.delete(`http://localhost:8080/api/siswa/${id}`)
        await Swal.fire({
          icon: 'success',
          title: 'Terhapus!',
          text: `Siswa ${siswa.nama} berhasil dihapus`,
          showConfirmButton: false,
          timer: 1500
        })
        fetchSiswa()
      } catch (error) {
        console.error('Gagal menghapus data:', error)
        Swal.fire({
          icon: 'error',
          title: 'Gagal',
          text: `Gagal menghapus siswa ${siswa.nama}`,
          confirmButtonColor: '#d33',
        })
      }
    }

    const submitForm = async () => {
      try {
        if (isEditing.value) {
          // Cari siswa asli untuk dibandingkan
          const originalSiswa = siswaList.value.find(s => s.id === formData.value.id)
          
          // Periksa apakah ada perubahan
          if (!hasChanges(originalSiswa, formData.value)) {
            await Swal.fire({
              icon: 'info',
              title: 'Tidak Ada Perubahan',
              text: `Data siswa ${originalSiswa.nama} tidak ada yang berubah`,
              confirmButtonColor: '#3085d6',
            })
            return
          }
          
          await axios.put(`http://localhost:8080/api/siswa/${formData.value.id}`, formData.value)
          await Swal.fire({
            icon: 'success',
            title: 'Berhasil!',
            text: `Data siswa ${formData.value.nama} berhasil diperbarui`,
            showConfirmButton: false,
            timer: 1500
          })
        } else {
          await axios.post("http://localhost:8080/api/siswa", formData.value)
          await Swal.fire({
            icon: 'success',
            title: 'Berhasil!',
            text: `Siswa ${formData.value.nama} berhasil ditambahkan`,
            showConfirmButton: false,
            timer: 1500
          })
        }
        fetchSiswa()
        closeForm()
      } catch (error) {
        console.error("Gagal menyimpan data:", error)
        await Swal.fire({
          icon: 'error',
          title: 'Gagal',
          text: isEditing.value 
            ? `Gagal memperbarui data siswa ${formData.value.nama}`
            : 'Gagal menambahkan siswa baru',
          confirmButtonColor: '#d33',
        })
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

    const jumlahPerKelas = computed(() => {
      const count = {}
      semuaSiswa.value.forEach((siswa) => {
        const key = `${siswa.jurusan} ${siswa.kelas}`
        count[key] = (count[key] || 0) + 1
      })

      const ordered = {}
      urutanKelas.forEach((kelas) => {
        ordered[kelas] = count[kelas] || 0
      })
      return ordered
    })

    onMounted(fetchSiswa)

    return {
      siswaList,
      semuaSiswa,
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
      jumlahPerKelas,
      fetchSiswa,
      addSiswa,
      editSiswa,
      deleteSiswa,
      submitForm,
      closeForm,
      resetFilter
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.title {
  text-align: center;
  margin-bottom: 30px;
  color: #333;
}

/* Filter Styling */
.filter-container {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  margin-bottom: 20px;
  align-items: flex-end;
}

.filter-group {
  flex: 1;
  min-width: 200px;
}

.filter-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 600;
  color: #555;
}

.filter-group select,
.filter-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.filter-group input:focus,
.filter-group select:focus {
  outline: none;
  border-color: #4a90e2;
  box-shadow: 0 0 0 2px rgba(74, 144, 226, 0.2);
}

.btn-refresh {
  background-color: #f0f0f0;
  color: #555;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 10px 15px;
  cursor: pointer;
  transition: all 0.3s;
  height: 42px;
}

.btn-refresh:hover {
  background-color: #e0e0e0;
}

/* Button Styling */
.btn-add, .btn-add-bulk {
  color: white;
  border: none;
  padding: 10px 15px;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 10px;
  margin-bottom: 20px;
  font-weight: 600;
  transition: background-color 0.3s;
}

.btn-add {
  background-color: #28a745;
}

.btn-add:hover {
  background-color: #218838;
}

.btn-add-bulk {
  background-color: #17a2b8;
}

.btn-add-bulk:hover {
  background-color: #138496;
}

/* Table + Rekap Container */
.table-container {
  display: flex;
  gap: 30px;
  position: relative;
  align-items: flex-start;
}

/* Tabel Utama */
.table-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  min-width: 0;
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  background: #fff;
  font-size: 14px;
}

thead {
  background-color: #343a40;
  color: #fff;
}

th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

tbody tr:hover {
  background-color: #f2f2f2;
}

.btn-edit, .btn-delete {
  padding: 6px 12px;
  margin-right: 5px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 13px;
  font-weight: 600;
}

.btn-edit {
  background-color: #ffc107;
  color: #000;
}

.btn-edit:hover {
  background-color: #e0a800;
}

.btn-delete {
  background-color: #dc3545;
  color: white;
}

.btn-delete:hover {
  background-color: #c82333;
}

/* Rekap Jumlah */
.rekap-jumlah {
  width: 30%;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  position: sticky;
  top: 80px;
  border: 1px solid #e0e0e0;
  display: flex;
  flex-direction: column;
  height: 70vh;
}

.rekap-header {
  padding: 15px 15px 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #f0f0f0;
}

.rekap-header h4 {
  margin: 0;
  font-size: 15px;
  color: #444;
  display: flex;
  align-items: center;
  gap: 8px;
}

.total-siswa {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #666;
}

.total-number {
  font-weight: 700;
  color: #27ab83;
  background-color: #e6f6f0;
  border-radius: 12px;
  padding: 2px 8px;
  font-size: 13px;
}

.rekap-scroll-container {
  overflow-y: auto;
  flex-grow: 1;
  padding: 10px 15px;
}

.rekap-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 10px;
}

.rekap-item {
  background-color: #f9f9f9;
  border-radius: 6px;
  padding: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: all 0.2s ease;
}

.rekap-item:hover {
  background-color: #f0f0f0;
  transform: translateY(-2px);
}

.kelas-label {
  font-size: 12px;
  color: #666;
  font-weight: 500;
  text-transform: uppercase;
}

.kelas-romawi {
  font-size: 14px;
  font-weight: 600;
  color: #333;
  margin: 3px 0;
}

.jumlah-siswa {
  font-size: 16px;
  font-weight: 700;
  color: #2c7be5;
  background-color: #f0f7ff;
  border-radius: 12px;
  padding: 2px 8px;
  min-width: 24px;
  text-align: center;
}

/* Modal Styling */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
}

.modal-content {
  background: white;
  padding: 30px;
  border-radius: 8px;
  width: 100%;
  max-width: 500px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.2);
}

.modal-content h3 {
  margin-bottom: 20px;
  font-size: 20px;
  color: #333;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  font-weight: 600;
  margin-bottom: 5px;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 10px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}

.btn-submit {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
}

.btn-submit:hover {
  background-color: #0069d9;
}

.btn-cancel {
  background-color: #6c757d;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
}

.btn-cancel:hover {
  background-color: #5a6268;
}

/* Responsive Tweaks */
@media (max-width: 768px) {
  .table-container {
    flex-direction: column;
  }

  .rekap-jumlah {
    width: 100%;
    position: static;
    top: auto;
    margin-top: 20px;
  }
}

/* SweetAlert Customization */
.swal2-popup {
  font-family: 'Arial', sans-serif;
}

.swal2-title {
  font-size: 1.2rem !important;
}

.swal2-confirm {
  padding: 8px 20px !important;
}
</style>