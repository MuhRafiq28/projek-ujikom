<template>
  <div class="landing-page pb">
    <h1>Welcome to TooPiket</h1>
    <p>Silakan login untuk melanjutkan.</p>
    <nuxt-link to="/login" class="login-btn">Login</nuxt-link>
  </div>
</template>

<script>
export default {
  name: 'LandingPage',
};
</script>

<style scoped>
.landing-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(135deg, #3498db, #8e44ad);
  color: white;
  text-align: center;
  font-family: 'Arial', sans-serif;
  padding: 20px;
  box-sizing: border-box;
}

h1 {
  font-size: 3.5em;
  font-weight: bold;
  color: #ffffff;
  margin-bottom: 20px;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.3);
  animation: fadeIn 2s ease-in-out;
}

p {
  font-size: 1.5em;
  color: #ffffff;
  margin-bottom: 30px;
  animation: fadeIn 3s ease-in-out;
}

.login-btn {
  padding: 15px 30px;
  font-size: 1.2em;
  background-color: #3498db;
  color: white;
  text-decoration: none;
  border-radius: 5px;
  transition: all 0.3s ease-in-out;
  display: inline-block;
  margin-top: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.login-btn:hover {
  background-color: #2980b9;
  transform: scale(1.1);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.3);
}

.login-btn:focus {
  outline: none;
  border: 2px solid #8e44ad;
}

@keyframes fadeIn {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
