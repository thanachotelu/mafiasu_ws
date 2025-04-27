<template>
  <div class="services-container">
    <h2 class="services-title">บริการของเรา</h2>
    <div class="services-grid">
      <router-link
        v-for="(service, index) in services"
        :key="index"
        :to="service.link"
        class="service-card-link"
      >
        <div class="service-card">
          <div class="service-image">
            <img :src="service.image" :alt="service.name" />
          </div>
          <h3 class="service-name">{{ service.name }}</h3>
          <p class="service-description">{{ service.description }}</p>
        </div>
      </router-link>
    </div>
  </div>
</template>

<script setup>
const services = [
  {
    name: 'เช่ารถระยะสั้น',
    description: 'รายวัน / รายสัปดาห์ / รายเดือน',
    image: '/src/assets/images/shortservice.png',
    link: '/short-rent-service',
  },
  {
    name: 'เช่ารถระยะยาว',
    description: 'เช่ารถยนต์ 3-5 ปี',
    image: '/src/assets/images/longservice.png',
    link: '/long-rent-service',
  },
  {
    name: 'เช่ารถสำหรับองค์กร',
    description: 'เช่ารถสำหรับบริษัทและองค์กร',
    image: '/src/assets/images/corporateservice.png',
    link: '/corporate-rent-service',
  },
  {
    name: 'เช่ารถพร้อมคนขับ',
    description: 'รถหรูพร้อมคนขับมืออาชีพ',
    image: '/src/assets/images/chaufferservice.png',
    link: '/chauffeur-service',
  },
];
</script>

<style scoped>
@import '../assets/main-service.css';

.service-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.service-card:hover .service-image img {
  transform: scale(1.05);
}
</style>