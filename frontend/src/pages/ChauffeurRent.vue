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
          <img src="../assets/images/car.png" alt="Icon" />
        </div>
        <h2 class="content-title">บริการรถเช่า พร้อมคนขับ</h2>
        <p class="content-description">
          พร้อมให้บริการรถเช่าพร้อมคนขับ สัญญารายเดือนและรายปี โดยพนักงานคนขับทีมงานมืออาชีพ พร้อมการบริการระดับสากล ให้คุณได้รับความสะดวกสบายตลอดการเดินทาง มั่นใจใช้บริการกับ บริษัท ไทยเร้นท์อะคาร์ คอร์ปอเรชั่น จำกัด 45 ปี แห่งผู้นำด้านตลาดรถเช่าในประเทศไทย ให้คุณได้มั่นใจกับการใช้บริการ โดยการบริการมาตรฐานระดับสากล
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