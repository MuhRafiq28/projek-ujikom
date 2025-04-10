<template>
  <div v-if="show" class="modal">
    <div class="modal-content">
      <div class="modal-header">
        <h2>{{ isEdit ? 'Edit Guru' : 'Tambah Guru' }}</h2>
        <button class="close-btn" @click="$emit('close')">&times;</button>
      </div>
      <form @submit.prevent="submitForm">
        <div class="form-group">
          <label for="nama">Nama</label>
          <input type="text" v-model="guruForm.nama" required />
        </div>
        <div class="form-group">
          <label for="kodeguru">Kode Guru</label>
          <input type="text" v-model="guruForm.kodeguru" required />
        </div>
        <div class="form-group">
          <label for="mapelguru">Mapel Guru</label>
          <input type="text" v-model="guruForm.mapelguru" required />
        </div>
        <div class="form-actions">
          <button type="button" class="btn-cancel" @click="$emit('close')">Batal</button>
          <button type="submit" class="btn-submit">{{ isEdit ? 'Update' : 'Tambah' }}</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    show: Boolean,
    isEdit: Boolean,
    guruData: Object
  },
  data() {
    return {
      guruForm: {
        id: null,
        nama: '',
        kodeguru: '',
        mapelguru: ''
      }
    };
  },
  watch: {
    guruData: {
      handler(newData) {
        if (newData) {
          this.guruForm = { ...newData };
        }
      },
      deep: true,
      immediate: true
    }
  },
  methods: {
    submitForm() {
    if (this.isEdit && !this.isFormChanged()) {this.$toast.error('Data tidak ada perubahan');


      return;
    }
    this.$emit('submit', this.guruForm);
  },
  isFormChanged() {
    return (
      this.guruForm.nama !== this.guruData.nama ||
      this.guruForm.kodeguru !== this.guruData.kodeguru ||
      this.guruForm.mapelguru !== this.guruData.mapelguru
    );
  },
  }
};
</script>

<style scoped>
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  backdrop-filter: blur(2px);
}

.modal-content {
  background-color: white;
  border-radius: 8px;
  width: 450px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
  overflow: hidden;
  animation: modalFadeIn 0.3s ease-out;
}

.modal-header {
  padding: 20px 24px;
  background-color: #4a6cf7;
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h2 {
  margin: 0;
  font-size: 1.4rem;
  font-weight: 600;
}

.close-btn {
  background: none;
  border: none;
  color: white;
  font-size: 1.5rem;
  cursor: pointer;
  padding: 0;
  line-height: 1;
  transition: transform 0.2s;
}

.close-btn:hover {
  transform: scale(1.1);
}

form {
  padding: 24px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #555;
}

.form-group input {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  transition: border-color 0.3s, box-shadow 0.3s;
}

.form-group input:focus {
  border-color: #4a6cf7;
  box-shadow: 0 0 0 2px rgba(74, 108, 247, 0.2);
  outline: none;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
}

button {
  padding: 10px 20px;
  border-radius: 4px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
}

.btn-submit {
  background-color: #4a6cf7;
  color: white;
}

.btn-submit:hover {
  background-color: #3a5ce4;
  transform: translateY(-1px);
}

.btn-cancel {
  background-color: #f5f5f5;
  color: #555;
}

.btn-cancel:hover {
  background-color: #e0e0e0;
}

@keyframes modalFadeIn {
  from {
      opacity: 0;
      transform: translateY(-20px);
  }
  to {
      opacity: 1;
      transform: translateY(0);
  }
}

@media (max-width: 480px) {
  .modal-content {
      width: 90%;
      margin: 0 auto;
  }
}
</style>
