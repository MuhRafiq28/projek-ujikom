<template>
  <div>
    <h2>Tambah Siswa</h2>
    <form @submit.prevent="tambahSiswa">
      <div>
        <label>NIS:</label>
        <input type="text" v-model="siswa.nis" required />
      </div>
      <div>
        <label>Nama:</label>
        <input type="text" v-model="siswa.nama" required />
      </div>
      <div>
        <label>Jurusan:</label>
        <input type="text" v-model="siswa.jurusan" required />
      </div>
      <div>
        <label>Kelas:</label>
        <input type="text" v-model="siswa.kelas" required />
      </div>
      <button type="submit">Simpan</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      siswa: {
        nis: "",
        nama: "",
        jurusan: "",
        kelas: ""
      }
    };
  },
  methods: {
    async tambahSiswa() {
      try {
        const response = await this.$axios.post("http://localhost:8080/siswa", this.siswa);
        alert(response.data.message);
        this.$router.push("/siswa"); // Redirect ke daftar siswa setelah sukses
      } catch (error) {
        alert("Gagal menambahkan siswa: " + error.response.data.error);
      }
    }
  }
};
</script>
