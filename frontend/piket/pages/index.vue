<template>
  <!-- From Uiverse.io by adamgiebl -->
  <section class="dots-container">
    <div class="dot"></div>
    <div class="dot"></div>
    <div class="dot"></div>
    <div class="dot"></div>
    <div class="dot"></div>
  </section>
</template>

<script>
export default {
  name: 'IndexPage',
  mounted() {
    const token = localStorage.getItem('authToken'); // Perbaiki key Token
    console.log("🔍 Token ditemukan:", token); // Debugging

    if (token) {
      this.$nextTick(() => {
        this.$router.push('/home');
      });
    }
  }
};
</script>

<style scoped>
/* Pusatkan animasi di tengah layar */
.dots-container {
  position: fixed; /* Pastikan elemen tetap di tengah saat scroll */
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: center;
  justify-content: center;
  height: auto;
  width: auto;
}

/* Animasi loading */
.dot {
  height: 20px;
  width: 20px;
  margin-right: 10px;
  border-radius: 10px;
  background-color: #b3d4fc;
  animation: pulse 1.5s infinite ease-in-out;
}

.dot:last-child {
  margin-right: 0;
}

.dot:nth-child(1) {
  animation-delay: -0.3s;
}

.dot:nth-child(2) {
  animation-delay: -0.1s;
}

.dot:nth-child(3) {
  animation-delay: 0.1s;
}

@keyframes pulse {
  0% {
    transform: scale(0.8);
    background-color: #b3d4fc;
    box-shadow: 0 0 0 0 rgba(178, 212, 252, 0.7);
  }

  50% {
    transform: scale(1.2);
    background-color: #6793fb;
    box-shadow: 0 0 0 10px rgba(178, 212, 252, 0);
  }

  100% {
    transform: scale(0.8);
    background-color: #b3d4fc;
    box-shadow: 0 0 0 0 rgba(178, 212, 252, 0.7);
  }
}
</style>
