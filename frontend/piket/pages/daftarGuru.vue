<template>
  <div class="container">
    <Navnew />

    <div class="content">
      <div class="header-section">
        <h1 class="title">Daftar Guru</h1>
        <div class="header-actions" v-if="userRole === 'admin'">
          <button @click="openModal" class="button add-button">
            <span class="button-icon">+</span> Tambah Guru
          </button>
        </div>
      </div>

      <div class="search-add-container">
        <div class="search-container">
          <input
            type="text"
            v-model="searchTerm"
            placeholder="Cari berdasarkan nama guru..."
            class="search-input"
          />
          <svg class="search-icon" viewBox="0 0 24 24">
            <path fill="currentColor" d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z" />
          </svg>
        </div>
      </div>

      <div class="table-responsive">
        <table class="table">
          <thead>
            <tr>
              <th class="text-center">No</th>
              <th>Nama</th>
              <th>Kode Guru</th>
              <th>Mapel Guru</th>
              <th v-if="userRole === 'admin'" class="text-center">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(guru, index) in filteredGurus" :key="guru.id" class="table-row">
              <td class="text-center">{{ index + 1 }}</td>
              <td>{{ guru.nama }}</td>
              <td>{{ guru.kodeguru }}</td>
              <td>{{ guru.mapelguru }}</td>
              <td class="action-buttons" v-if="userRole === 'admin'">
                <div class="flex justify-center space-x-2">
                  <button @click="editGuru(guru)" class="button edit-button">
                    <span class="button-icon">✏️</span> Edit
                  </button>
                  <button @click="deleteGuru(guru)" class="button delete-button">
                    <span class="button-icon">🗑</span> Hapus
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
        <div v-if="filteredGurus.length === 0" class="empty-state">
          <svg class="empty-icon" viewBox="0 0 24 24">
            <path fill="currentColor" d="M11,15H13V17H11V15M11,7H13V13H11V7M12,2C6.47,2 2,6.5 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4A8,8 0 0,1 20,12A8,8 0 0,1 12,20Z" />
          </svg>
          <p class="empty-text">Data guru tidak ditemukan</p>
        </div>
      </div>

      <ModalGuru
        :show="showModal"
        :isEdit="isEdit"
        :guruData="guruForm"
        @submit="handleSubmit"
        @close="closeModal"
      />
    </div>
  </div>
</template>

<script>
import ModalGuru from '@/components/ModalGuru.vue';
import Navnew from '@/components/Navnew.vue';
import Swal from 'sweetalert2';

export default {
  components: { ModalGuru, Navnew },
  data() {
    return {
      gurus: [],
      showModal: false,
      isEdit: false,
      userRole: 'user',
      guruForm: { id: null, nama: '', kodeguru: '', mapelguru: '' },
      searchTerm: ''
    };
  },
  computed: {
    filteredGurus() {
      return this.gurus.filter(guru =>
        guru.nama.toLowerCase().includes(this.searchTerm.toLowerCase())
      );
    }
  },
  async mounted() {
    if (process.client) {
      this.userRole = localStorage.getItem('userRole') || 'user';
    }
    await this.fetchGurus();
  },
  methods: {
    async fetchGurus() {
      try {
        const response = await this.$axios.get('http://localhost:8080/api/guru');
        this.gurus = response.data;
      } catch (error) {
        console.error('Gagal memuat data guru:', error);
        Swal.fire({
          icon: 'error',
          title: 'Gagal!',
          text: 'Gagal memuat data guru',
          showConfirmButton: false,
          timer: 1500
        });
      }
    },
    openModal() {
      this.showModal = true;
      this.isEdit = false;
      this.guruForm = { id: null, nama: '', kodeguru: '', mapelguru: '' };
    },
    closeModal() {
      this.showModal = false;
    },
    editGuru(guru) {
      this.isEdit = true;
      this.guruForm = { ...guru };
      this.showModal = true;
    },
    async handleSubmit(guruData) {
      try {
        if (this.isEdit) {
          await this.updateGuru(guruData);
        } else {
          await this.createGuru(guruData);
        }
        this.closeModal();
        await this.fetchGurus();
      } catch (error) {
        console.error('Gagal menyimpan data:', error);
        Swal.fire({
          icon: 'error',
          title: 'Gagal!',
          text: 'Gagal menyimpan data guru',
          showConfirmButton: false,
          timer: 1500
        });
      }
    },
    async createGuru(guruData) {
      try {
        await this.$axios.post('http://localhost:8080/api/guru', guruData);
        Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: `Data guru ${guruData.nama} berhasil ditambahkan`,
          showConfirmButton: false,
          timer: 1500
        });
      } catch (error) {
        throw error;
      }
    },
    async updateGuru(guruData) {
      try {
        await this.$axios.put(`http://localhost:8080/api/guru/${guruData.id}`, guruData);
        Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: `Data guru ${guruData.nama} berhasil diperbarui`,
          showConfirmButton: false,
          timer: 1500
        });
      } catch (error) {
        throw error;
      }
    },
    async deleteGuru(guru) {
      const result = await Swal.fire({
        title: 'Apakah Anda yakin?',
        text: `Anda akan menghapus data guru ${guru.nama}`,
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Ya, hapus!',
        cancelButtonText: 'Batal'
      });

      if (result.isConfirmed) {
        try {
          await this.$axios.delete(`http://localhost:8080/api/guru/${guru.id}`);
          await this.fetchGurus();
          Swal.fire({
            icon: 'success',
            title: 'Terhapus!',
            text: `Data guru ${guru.nama} telah dihapus`,
            showConfirmButton: false,
            timer: 1500
          });
        } catch (error) {
          console.error('Gagal menghapus guru:', error);
          Swal.fire({
            icon: 'error',
            title: 'Gagal!',
            text: 'Gagal menghapus data guru',
            showConfirmButton: false,
            timer: 1500
          });
        }
      }
    }
  }
};
</script>

<style scoped>
/* CSS tetap sama seperti sebelumnya */
.container {
  width: 100%;
  min-height: 100vh;
  background-color: #f8fafc;
}

.content {
  max-width: 1200px;
  margin: 80px auto 30px;
  padding: 30px;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.title {
  font-size: 1.75rem;
  font-weight: 700;
  color: #1e293b;
  margin: 0;
}

.search-add-container {
  display: flex;
  justify-content: space-between;
  margin-bottom: 24px;
}

.search-container {
  position: relative;
  width: 300px;
}

.search-input {
  width: 100%;
  padding: 10px 15px 10px 40px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 0.95rem;
  transition: all 0.3s ease;
  background-color: #f8fafc;
}

.search-input:focus {
  outline: none;
  border-color: #6366f1;
  box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
  background-color: white;
}

.search-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  width: 18px;
  height: 18px;
  color: #94a3b8;
}

.table-responsive {
  overflow-x: auto;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
}

.table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.95rem;
}

th {
  background-color: #4f46e5;
  color: white;
  font-weight: 600;
  padding: 12px 16px;
  text-align: left;
}

th.text-center {
  text-align: center;
}

td {
  padding: 12px 16px;
  border-bottom: 1px solid #e2e8f0;
  color: #334155;
}

td.text-center {
  text-align: center;
}

.table-row:hover {
  background-color: #f8fafc;
}

.button {
  display: inline-flex;
  align-items: center;
  padding: 8px 16px;
  border-radius: 6px;
  font-weight: 500;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
  outline: none;
}

.button-icon {
  margin-right: 6px;
}

.add-button {
  background-color: #4f46e5;
  color: white;
}

.add-button:hover {
  background-color: #4338ca;
  transform: translateY(-1px);
}

.edit-button {
  background-color: #f59e0b;
  color: white;
}

.edit-button:hover {
  background-color: #d97706;
}

.delete-button {
  background-color: #ef4444;
  color: white;
}

.delete-button:hover {
  background-color: #dc2626;
}

.empty-state {
  padding: 40px 20px;
  text-align: center;
  background-color: #f8fafc;
}

.empty-icon {
  width: 48px;
  height: 48px;
  margin-bottom: 16px;
  color: #94a3b8;
}

.empty-text {
  color: #64748b;
  font-size: 1rem;
  margin: 0;
}

@media (max-width: 768px) {
  .content {
    padding: 20px;
    margin-top: 70px;
  }

  .header-section {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .search-add-container {
    flex-direction: column;
    gap: 16px;
  }

  .search-container {
    width: 100%;
  }

  .action-buttons {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }
}
</style>
