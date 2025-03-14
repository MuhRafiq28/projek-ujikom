<template>
  <div class="container">
    <Navnew />

    <div class="content ">
      <h1 class="title">Daftar Guru</h1>
      <div class="button-container" v-if="userRole === 'admin'">
        <h1 class="title">Daftar Guru</h1>
        <button @click="openModal" class="button add-button">+ Tambah Guru</button>
      </div>

      <div class="table-container mt-2">
        <table class="table">
          <thead>
            <tr>
              <th>No</th>
              <th>Nama</th>
              <th>Kode Guru</th>
              <th>Mapel Guru</th>
              <th v-if="userRole === 'admin'">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(guru, index) in gurus" :key="guru.id" class="table-row">
              <td>{{ index + 1 }}</td>
              <td>{{ guru.nama }}</td>
              <td>{{ guru.kodeguru }}</td>
              <td>{{ guru.mapelguru }}</td>
              <td class="action-buttons" v-if="userRole === 'admin'">
                <button @click="editGuru(guru)" class="button edit-button">✏️ Edit</button>
                <button @click="deleteGuru(guru.id)" class="button delete-button mt-1">🗑 Hapus</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <p v-if="gurus.length === 0" class="no-data">Data guru tidak ditemukan.</p>

      <ModalGuru :show="showModal" :isEdit="isEdit" :guruData="guruForm" @submit="handleSubmit" @close="closeModal" />
    </div>
  </div>
</template>

<script>
import ModalGuru from '@/components/ModalGuru.vue';
import Navnew from '@/components/Navnew.vue';

export default {
  components: { ModalGuru, Navnew },
  data() {
    return {
      gurus: [],
      showModal: false,
      isEdit: false,
      userRole: 'user', // Default value
      guruForm: { id: null, nama: '', kodeguru: '', mapelguru: '' }
    };
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
        console.error('Error fetching gurus:', error);
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
      if (this.isEdit) {
        await this.updateGuru(guruData);
      } else {
        await this.createGuru(guruData);
      }
      this.closeModal();
      await this.fetchGurus();
    },
    async createGuru(guruData) {
      try {
        await this.$axios.post('http://localhost:8080/api/guru', guruData);
      } catch (error) {
        console.error('Error creating guru:', error);
      }
    },
    async updateGuru(guruData) {
      try {
        await this.$axios.put(`http://localhost:8080/api/guru/${guruData.id}`, guruData);
      } catch (error) {
        console.error('Error updating guru:', error);
      }
    },
    async deleteGuru(id) {
      try {
        await this.$axios.delete(`http://localhost:8080/api/guru/${id}`);
        await this.fetchGurus();
      } catch (error) {
        console.error('Error deleting guru:', error);
      }
    }
  }
};
</script>

<style scoped>
/* Membuat container memenuhi layar */
.container {
  width: 100%;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}

/* Membuat konten memenuhi layar */
.content {
  width: 100%;
  margin-top: 60px;
  max-width: none;
  background-color: #f9fafb;
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.button-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  padding: 0 10px;
}

.title {
  font-size: 2rem;
  font-weight: 600;
  color: #333;
}

/* Tombol */
.button {
  padding: 10px 20px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  border: none;
  outline: none;
  font-size: 1rem;
}

.add-button {
  background-color: #3490dc;
  color: white;
}

.add-button:hover {
  background-color: #2779bd;
}

/* Table dibuat penuh */
.table-container {
  width: 100%;
  overflow-x: auto;
}

.table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;
}

th, td {
  padding: 12px 20px;
  text-align: center; /* Rata tengah secara horizontal */
  vertical-align: middle; /* Rata tengah secara vertikal */
  font-size: 1rem;
}

th {
  background-color: #3800bb;
  color: white;
}

.table-row:nth-child(even) {
  background-color: #f8fafc;
}

.table-row:hover {
  background-color: #f1f5f9;
}

/* Pesan jika data kosong */
.no-data {
  text-align: center;
  color: #6b7280;
  font-size: 1.125rem;
  margin-top: 20px;
}
</style>
