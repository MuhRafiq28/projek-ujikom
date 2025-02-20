<template>
  <div>
    <Navbar />
  <div class="home-user">
    <h1>Welcome to Home Admin Page</h1>
    <!-- Add your content here -->
  </div>
</div>
</template>

<script>
import Navbar from '~/components/Navbar.vue';

export default {
  name: 'HomeUser',
  comments: {
    Navbar
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
