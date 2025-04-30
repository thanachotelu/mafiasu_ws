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
          <img src="../assets/images/calendar.png" alt="Icon" />
        </div>
        <h2 class="content-title">รถเช่าระยะยาว</h2>
        <p class="content-description">
          บริการรถเช่าระยะยาว รถเช่าขับเอง (รายปี) มาตรฐานสากล จากไทยเร้นท์อะคาร์ การันตีคุณภาพรถ และการบริการโดยทีมงานมืออาชีพ ไม่ว่าจะเช่ารถสำหรับการท่องเที่ยว เช่ารถเพื่อใช้งานในชีวิตประจำวัน หรือ รถเช่าระหว่างซ่อม 
        </p>
      </div>
    </div>

    <Services :serviceList="longRentCars" />
  </div>
</template>

<script setup>
import Services from '../components/Services.vue';

const longRentCars = [
{
    brand: 'Isuzu',
    model: 'All New Mu-X',
    rentalPricePerDay: 0,
    seat: 7,
    doors: 5,
    gearType: 'อัตโนมัติ 6 สปีด',
    fuelType: 'ดีเซล',
    image: '/src/assets/images/All New Mu-X.png'
  },
  {
    brand: 'Isuzu',
    model: 'D-Max Cab4 1.9AT',
    rentalPricePerDay: 0,
    seat: 5,
    doors: 4,
    gearType: 'อัตโนมัติ CVT',
    fuelType: 'ดีเซล',
    image: '/src/assets/images/D-Max Cab4 1.9AT.png'
  },
  {
    brand: 'Toyota',
    model: 'Hiace',
    rentalPricePerDay: 0,
    seat: 12,
    doors: 4,
    gearType: 'อัตโนมัติ',
    fuelType: 'เบนซิน',
    image: '/src/assets/images/Hiace.png'
  }
];
</script>

<style scoped>
@import '../assets/rent-service.css';

</style>