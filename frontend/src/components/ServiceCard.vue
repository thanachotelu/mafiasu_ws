<template>
    <div class="service-card">
      <div class="service-image">
        <!-- Placeholder for service image -->
      </div>
      <div class="service-content">
        <h2>{{ service.name }}</h2>
        <p>{{ service.description }}</p>
        <p>{{ service.details }}</p>
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