<template>
  <transition name="fade">
    <div v-if="isVisible" class="modal-overlay">
      <div class="modal-container">
        <h2 class="modal-title">
          {{ editMode ? "Edit Jadwal" : "Tambah Jadwal" }}
        </h2>

        <form @submit.prevent="submit">
          <div class="form-group">
            <label>Hari:</label>
            <input v-model="localSchedule.day" type="text" placeholder="Contoh: Senin" required />
          </div>

          <div class="form-group">
            <label>Guru:</label>
            <input v-model="localSchedule.Guru" type="text" placeholder="Nama Guru" required />
          </div>

          <div class="form-group">
            <label>Waktu:</label>
            <input v-model="localSchedule.time" type="text" placeholder="08:00 - 10:00" required />
          </div>

          <div class="form-group">
            <label>Catatan:</label>
            <textarea v-model="localSchedule.notes" placeholder="Tambahkan catatan..." required></textarea>
          </div>

          <div class="button-group">
            <button type="submit" class="btn-save">
              {{ editMode ? "Update" : "Tambah" }}
            </button>
            <button type="button" class="btn-cancel" @click="$emit('close')">Batal</button>
          </div>
        </form>
      </div>
    </div>
  </transition>
</template>

<script>
import Swal from 'sweetalert2'

export default {
  props: {
    isVisible: Boolean,
    editMode: Boolean,
    schedule: Object
  },
  data() {
    return {
      localSchedule: { ...this.schedule }
    };
  },
  watch: {
    schedule(newVal) {
      this.localSchedule = { ...newVal };
    }
  },
  methods: {
    async submit() {
      try {
        // Emit event submit ke parent component
        this.$emit("submit", this.localSchedule);

        // Tampilkan notifikasi sukses
        await Swal.fire({
          title: 'Sukses!',
          text: this.editMode ? 'Jadwal berhasil diperbarui!' : 'Jadwal berhasil ditambahkan!',
          icon: 'success',
          confirmButtonText: 'OK',
          confirmButtonColor: '#6c63ff',
          timer: 3000,
          timerProgressBar: true,
          toast: true,
          position: 'top-end',
          showConfirmButton: false
        });

        // Tutup modal setelah notifikasi
        this.$emit('close');
      } catch (error) {
        // Tampilkan notifikasi error jika ada masalah
        Swal.fire({
          title: 'Error!',
          text: 'Terjadi kesalahan saat menyimpan data',
          icon: 'error',
          confirmButtonText: 'OK',
          confirmButtonColor: '#ff6b6b'
        });
      }
    }
  }
};
</script>

<style scoped>
/* Animasi Modal */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}

/* Overlay Modal */
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
  z-index: 999;
}

/* Container Modal */
.modal-container {
  background: linear-gradient(135deg, #fefefe, #f4f4f4);
  padding: 20px;
  border-radius: 12px;
  width: 420px;
  text-align: center;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.15);
  animation: scaleUp 0.3s ease-in-out;
}

@keyframes scaleUp {
  from {
    transform: scale(0.85);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}

/* Judul Modal */
.modal-title {
  font-size: 1.6rem;
  font-weight: 600;
  color: #444;
  margin-bottom: 15px;
}

/* Form Group */
.form-group {
  text-align: left;
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  font-size: 0.9rem;
  font-weight: 600;
  color: #555;
  margin-bottom: 5px;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 6px;
  font-size: 1rem;
  background-color: #f9f9f9;
  transition: 0.2s ease-in-out;
}

.form-group input:focus,
.form-group textarea:focus {
  border-color: #6c63ff;
  outline: none;
  background-color: #fff;
  box-shadow: 0px 0px 8px rgba(108, 99, 255, 0.3);
}

/* Tombol */
.button-group {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.btn-save {
  background: #6c63ff;
  color: white;
  padding: 10px 18px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: 0.3s ease;
}

.btn-save:hover {
  background: #554ef0;
}

.btn-cancel {
  background: #ff6b6b;
  color: white;
  padding: 10px 18px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  transition: 0.3s ease;
}

.btn-cancel:hover {
  background: #e45050;
}
</style>
