<template>
  <div>
    <NavAdmin />
    <div class="home-user">
      <h1>Welcome to Home Petugas Page</h1>
      <!-- Add your content here -->
    </div>
  </div>
</template>

<script>
import NavAdmin from '../components/NavAdmin.vue'; // Pastikan path ini benar

export default {
  name: 'HomeUser',
  components: {
    NavAdmin, // Menggunakan components dengan benar
  },
  data() {
    return {
      // Add your data properties here
    };
  },
  methods: {
    // Add your methods here
  },
  mounted() {
    // Add your mounted logic here
  }
};
</script>

<style scoped>
.home-user {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  font-size: 24px;
  color: #333;
}
</style>
