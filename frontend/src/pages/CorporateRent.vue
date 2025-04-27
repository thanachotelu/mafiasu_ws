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
          <img src="../assets/images/building.png" alt="Icon" />
        </div>
        <h2 class="content-title">รถเช่าสำหรับองค์กร</h2>
        <p class="content-description">
          ผู้นำธุรกิจบริการรถยนต์เช่าเพื่อองค์กร สามารถจัดหารถทุกประเภทตามความต้องการใช้งานและรูปแบบที่ลูกค้าต้องการ ด้วยอัตราค่าเช่าคงที่ตลอดอายุสัญญาการเช่า มีให้บริการทั้งเช่ารถเปล่า และเช่ารถพร้อมคนขับให้เลือกตาม ความเหมาะสมโดยอัตราค่าเช่าต่อเดือน จะรวมการบริการเรื่องภาษีประจำปี การจด/ต่อทะเบียน ค่าประกันภัยชั้น 1 ค่าบำรุงรักษาต่างๆและมีรถสำรองทดแทนพร้อมให้บริการภายใน 24 ชั่วโมง (เป็นการบริหารต้นทุน คุ้มค่า)
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