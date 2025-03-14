<template>
  <div>
    <Navnew />
    <div class="schedule-container">
      <h2>Semua Jadwal</h2>

      <!-- Menampilkan jadwal dengan maksimal dua kartu per baris -->
      <div class="schedule-grid" :class="{'single-item': schedules.length === 1}">
        <div v-for="(schedule, index) in schedules" :key="schedule.id"
             class="schedule-card"
             :class="{'new-row': index % 2 === 0}">
          <strong class="day-label">{{ schedule.day }}</strong>
          <p class="time-label">{{ schedule.time }}</p>
          <p><strong>Guru:</strong> {{ schedule.Guru }}</p>
          <p><strong>Catatan:</strong> {{ schedule.notes }}</p>

          <!-- Tombol Edit & Hapus hanya muncul jika user adalah admin -->
          <div v-if="userRole === 'admin'" class="action-buttons">
            <button class="edit-button" @click="prepareEdit(schedule)">✏️ Edit</button>

          </div>
        </div>
      </div>

      <!-- Modal CRUD -->
      <ModalJadwal
        :isVisible="modalOpen"
        :schedule="schedule"
        :editMode="editMode"
        @close="closeModal"
        @submit="handleSubmit"
      />
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ModalJadwal from '@/components/ModalJadwal.vue';
import Navnew from '@/components/Navnew.vue';

export default {
  components: {
    ModalJadwal,
    Navnew,
  },
  data() {
    return {
      schedules: [],
      schedule: { id: null, day: "", Guru: "", time: "", notes: "" },
      editMode: false,
      modalOpen: false,
      userRole: '', // Menyimpan peran pengguna dari localStorage
    };
  },
  mounted() {
    this.checkUserRole(); // Pastikan role diambil saat komponen dimuat
    this.fetchSchedules();
  },
  methods: {
    checkUserRole() {
      const userRole = localStorage.getItem('userRole');
      this.userRole = userRole || 'user';
      console.log('User Role:', this.userRole);
    },
    async fetchSchedules() {
      try {
        const response = await axios.get("http://localhost:8080/api/schedules");
        this.schedules = response.data;
      } catch (error) {
        console.error("Error fetching schedules:", error);
      }
    },
    prepareAdd() {
      this.schedule = { id: null, day: "", Guru: "", time: "", notes: "" };
      this.editMode = false;
      this.modalOpen = true;
    },
    prepareEdit(schedule) {
      this.schedule = { ...schedule };
      this.editMode = true;
      this.modalOpen = true;
    },
    closeModal() {
      this.modalOpen = false;
    },
    async handleSubmit(schedule) {
      if (this.editMode) {
        await this.updateSchedule(schedule);
      } else {
        await this.addSchedule(schedule);
      }
      this.closeModal();
    },
    async addSchedule(schedule) {
      try {
        const response = await axios.post("http://localhost:8080/api/schedules", schedule);
        this.schedules.push(response.data);
      } catch (error) {
        console.error("Error adding schedule:", error);
      }
    },
    async updateSchedule(schedule) {
      try {
        const response = await axios.put(`http://localhost:8080/api/schedules/${schedule.id}`, schedule);
        const index = this.schedules.findIndex((s) => s.id === schedule.id);
        if (index !== -1) {
          this.schedules.splice(index, 1, response.data);
        }
      } catch (error) {
        console.error("Error updating schedule:", error);
      }
    },
    async deleteSchedule(id) {
      try {
        await axios.delete(`http://localhost:8080/api/schedules/${id}`);
        this.schedules = this.schedules.filter(schedule => schedule.id !== id);
      } catch (error) {
        console.error("Error deleting schedule:", error);
      }
    },
    async mounted() {
    await this.fetchGurus();
    this.userRole = localStorage.getItem('userRole') || 'user';
 // Ambil peran pengguna dari localStorage
  },
  }
};
</script>

<style scoped>
.schedule-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  margin-top: 60px;
}

h2 {
  font-size: 1.8rem;
  color: #333;
  text-align: center;
  margin-bottom: 20px;
}

.add-button {
  display: block;
  margin: 0 auto 20px;
  padding: 12px 20px;
  background-color: #5c00cb;
  color: white;
  font-weight: bold;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.3s, transform 0.2s;
}

.add-button:hover {
  background-color: #3366ff;
  transform: scale(1.05);
}

.schedule-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  justify-content: center;
  color: white;
  justify-content: space-between;
}

.schedule-card {
  background: linear-gradient(
    135deg,
    #3498db,
    #8e44ad
  );
  padding: 20px;
  border-radius: 10px;
  text-align: center;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.2);
  transition: transform 0.3s;
}

.schedule-card:hover {
  transform: translateY(-5px);
}

.edit-button {
  background-color: #00d0ff;
  color: white;
  border: none;
  padding: 8px 15px;
  border-radius: 6px;
  cursor: pointer;
  transition: background 0.3s;
}

.edit-button:hover {
  background-color: #e68900;
}
</style>
