<template>
  <div class="container mt-5">
    <h3 class="text-center mb-3">Jadwal Siswa</h3>
    <div v-if="jadwal.length > 0">
      <div v-for="item in jadwal" :key="item.id" class="card mb-2 p-3 shadow-sm">
        <h5>{{ item.hari }} - {{ item.jam_mulai }} - {{ item.jam_selesai }}</h5>
        <p class="mb-1"><strong>Mata Pelajaran:</strong> {{ item.mata_pelajaran }}</p>
        <p class="mb-1"><strong>Guru:</strong> {{ item.guru }}</p>
      </div>
    </div>
    <p v-else class="text-center text-muted">Jadwal tidak ditemukan.</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      jadwal: [],
    };
  },
  mounted() {
    this.getJadwal();
  },
  methods: {
    async getJadwal() {
      try {
        const res = await this.$axios.get("http://localhost:8080/api/jadwal");
        this.jadwal = res.data || [];
      } catch (err) {
        console.error("Gagal mengambil jadwal:", err);
      }
    },
  },
};
</script>

<style>
.container {
  max-width: 400px;
}
.card {
  border-radius: 10px;
}
</style>
