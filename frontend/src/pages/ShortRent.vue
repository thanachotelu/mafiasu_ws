<template>
  <div class="short-rent container">
    <div class="top-banner">
      <div class="banner-content">
        <h1 class="banner-title">ค้นหารถเช่า<br />ที่เหมาะกับคุณ</h1>
        <p class="banner-subtitle">
          กรองผลลัพธ์ตามความต้องการของคุณ<br />
          เรามีรถหลากหลายรุ่นให้เลือก
        </p>
        <router-link to="/car-lists" class="banner-button">ดูเพิ่มเติม</router-link>

      </div>
    </div>

    <div class="content">
      <div class="content-header">
        <div class="content-icon">
          <img src="../assets/images/clock.png" alt="Icon" />
        </div>
        <h2 class="content-title">รถเช่าระยะสั้น</h2>
        <p class="content-description">
          บริการรถเช่าระยะสั้น รถเช่าขับเอง (รายวัน-รายสัปดาห์-รายเดือน) มาตรฐานสากล
          จากไทยเร้นท์อะคาร์ การันตีคุณภาพรถ และการบริการ
          โดยทีมงานมืออาชีพ ไม่ว่าจะเช่ารถสำหรับการท่องเที่ยว เช่ารถเพื่อใช้งานในชีวิตประจำวัน
          หรือ รถเช่าระหว่างซ่อม
        </p>
      </div>
    </div>

    <Services />
  </div>
</template>

<script setup>
import Services from '../components/Services.vue';
</script>

<style scoped>
@import '../assets/rent-service.css';

</style>