<template>
    <div v-if="show" class="modal">
      <div class="modal-content">
        <h2>{{ isEdit ? 'Edit Guru' : 'Tambah Guru' }}</h2>
        <form @submit.prevent="submitForm">
          <div>
            <label for="nama">Nama</label>
            <input type="text" v-model="guruForm.nama" required />
          </div>
          <div>
            <label for="kodeguru">Kode Guru</label>
            <input type="text" v-model="guruForm.kodeguru" required />
          </div>
          <div>
            <label for="mapelguru">Mapel Guru</label>
            <input type="text" v-model="guruForm.mapelguru" required />
          </div>
          <button type="submit">{{ isEdit ? 'Update' : 'Tambah' }}</button>
          <button type="button" @click="$emit('close')">Batal</button>
        </form>
      </div>
    </div>
  </template>

  <script>
  export default {
    props: {
      show: Boolean, // Menentukan apakah modal ditampilkan
      isEdit: Boolean, // Apakah sedang dalam mode edit
      guruData: Object // Data guru untuk diedit
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
        this.$emit('submit', this.guruForm);
      }
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
  }
  .modal-content {
    background-color: white;
    padding: 20px;
    border-radius: 5px;
    width: 400px;
  }
  button {
    padding: 10px 20px;
    margin: 10px;
  }
  </style>
