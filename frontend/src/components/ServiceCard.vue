<template>
  <div class="service-card">
    <div class="service-image">
      <img :src="service.image" :alt="`${service.brand} ${service.model}`" />
    </div>
    <div class="service-content">
      <h2>{{ service.brand }} {{ service.model }}</h2>
      <p>ที่นั่ง: {{ service.seat }} | ประตู: {{ service.doors }}</p>
      <p>เกียร์: {{ service.gearType }} | เชื้อเพลิง: {{ service.fuelType }}</p>
      <p class="price">{{ service.rentalPricePerDay.toFixed(2) }} บาท/วัน</p>
      <button class="read-more-btn">Read More</button>
    </div>
  </div>
</template>

<script setup>
defineProps({
  service: {
    type: Object,
    required: true
  }
});
</script>

<style scoped>
@import '../assets/service-card.css';
</style>
