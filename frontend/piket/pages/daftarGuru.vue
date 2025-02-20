<template>
  <div>
  <Navbar />
  <div class="container">
    <h1 class="title">Daftar Guru</h1>

    <!-- Tombol Tambah Guru (Hanya untuk Admin) -->
    <div class="button-container" v-if="userRole === 'admin'">
      <button @click="openModal" class="button add-button">
        + Tambah Guru
      </button>
    </div>

    <!-- Tabel Daftar Guru -->
    <div class="table-container">
      <table class="table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Nama</th>
            <th>Kode Guru</th>
            <th>Mapel Guru</th>
            <th v-if="userRole === 'admin'">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="guru in gurus" :key="guru.id" class="table-row">
            <td>{{ guru.id }}</td>
            <td>{{ guru.nama }}</td>
            <td>{{ guru.kodeguru }}</td>
            <td>{{ guru.mapelguru }}</td>
            <td class="action-buttons" v-if="userRole === 'admin'">
              <button @click="editGuru(guru)" class="button edit-button">
                ✏️ Edit
              </button>
              <button @click="deleteGuru(guru.id)" class="button delete-button">
                🗑 Hapus
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <p v-if="gurus.length === 0" class="no-data">Data guru tidak ditemukan.</p>

    <!-- Modal CRUD -->
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
import Navbar from '@/components/Navbar.vue';
export default {
  components: { ModalGuru, Navbar },
  data() {
    return {
      gurus: [],
      showModal: false,
      isEdit: false,
      userRole: '', // Menyimpan peran pengguna dari localStorage
      guruForm: {
        id: null,
        nama: '',
        kodeguru: '',
        mapelguru: ''
      }
    };
  },
  async mounted() {
    await this.fetchGurus();
    this.userRole = localStorage.getItem('userRole') || 'user';
 // Ambil peran pengguna dari localStorage
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
        console.log('Guru berhasil ditambahkan');
      } catch (error) {
        console.error('Error creating guru:', error);
      }
    },
    async updateGuru(guruData) {
      try {
        await this.$axios.put(`http://localhost:8080/api/guru/${guruData.id}`, guruData);
        console.log('Guru berhasil diperbarui');
      } catch (error) {
        console.error('Error updating guru:', error);
      }
    },
    async deleteGuru(id) {
      try {
        await this.$axios.delete(`http://localhost:8080/api/guru/${id}`);
        console.log('Guru berhasil dihapus');
        await this.fetchGurus();
      } catch (error) {
        console.error('Error deleting guru:', error);
      }
    }
  }
};
</script>

<style scoped>
template {
  background: #757171;
}
.container {
  max-width: 900px;
  margin: 0 auto;
  padding: 24px;
  background-color: #f9fafb;
  border-radius: 12px;
  margin-top: 60px;
}

.title {
  font-size: 2.25rem;
  font-weight: 600;
  text-align: center;
  color: #333;
  margin-bottom: 32px;
}

.button-container {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 24px;
}

.button {
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  border: none;
  outline: none;
  font-size: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.add-button {
  background-color: #3490dc;
  color: white;
}

.add-button:hover {
  background-color: #2779bd;
  transform: translateY(-2px);
}

.table-container {
  overflow-x: auto;
  background-color: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  margin-top: 24px;
}

.table {
  width: 100%;
  border-collapse: collapse;
  border-radius: 8px;
  overflow: hidden;
}

th,
td {
  padding: 16px 24px;
  text-align: left;
  color: #333;
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

.action-buttons {
  text-align: center;
}

.edit-button {
  background: linear-gradient(
    135deg,
    #3498db,
    #8e44ad
  );
  color: white;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
}

.edit-button:hover {
  background-color: #f59e0b;
  transform: translateY(-2px);
}

.delete-button {
  background: linear-gradient(
    135deg,
    #3498db,
    #8e44ad
  );
  color: white;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  margin-left: 8px;
}

.delete-button:hover {
  background-color: #dc2626;
  transform: translateY(-2px);
}

.no-data {
  text-align: center;
  color: #6b7280;
  font-size: 1.125rem;
  margin-top: 32px;
}
</style>
