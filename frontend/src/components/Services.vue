<template>
  <div class="services-section">
    <ServiceCard
      v-for="(service, index) in serviceList"
      :key="index"
      :service="service"
    />
  </div>
</template>

<script setup>
import { defineProps } from 'vue';
import ServiceCard from './ServiceCard.vue';

const props = defineProps({
  serviceList: {
    type: Array,
    required: true,
  }
});
</script>

<style scoped>
@import '../assets/services.css';
</style>