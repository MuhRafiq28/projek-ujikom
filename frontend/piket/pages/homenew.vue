<template>
  <div>
    <Navnew />
    <div class="content">
      <div class="text-content">
        <h1>INTERNASIONAL ACADEMY</h1>
        <p>
          Selamat datang di web piket SMK kami. Kerjakan tugas Anda dengan santai agar tidak terasa berat.
        </p>
      </div>
      <div class="image-hero">
        <img src="/images/inspirasi-logo.png" alt="SMKN 1 Cisarua" />
      </div>
    </div>
    <footer class="footer">
      <p>&copy; 2025 Muhammad Rafiq. All Rights Reserved.</p>
    </footer>
  </div>
</template>

<script>
import Navnew from "../components/Navnew.vue";
export default {
  components: {
    Navnew,
  },
};
</script>

<style scoped>
.content {
  margin-top: 55px;
  background-image: url("/images/hero-images.jpg");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  height: 473px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 50px;
  padding: 0 100px;
  color: white;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5);
}

.text-content {
  flex: 1;
  text-align: left;
  max-width: 500px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  margin-left: 50px;
}

.text-content h1 {
  font-size: 36px;
  font-weight: bold;
  margin-bottom: 10px;
}

.text-content p {
  font-size: 18px;
  line-height: 1.5;
}

.image-hero {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.image-hero img {
  width: 100%;
  max-width: 300px;
  height: auto;
  object-fit: cover;
}

.footer {
  background: #96858F;
  color: white;
  text-align: center;
  padding: 15px 0;
  font-size: 14px;
  position: relative;
  bottom: 0;
  width: 100%;
}
</style>
