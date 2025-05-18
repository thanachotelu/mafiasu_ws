<template>
  <NavBar1 />
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

    <Services :serviceList="shortRentCars" />
  </div>
</template>

<script setup>
import NavBar1 from '../components/NavBar1.vue';
import Services from '../components/Services.vue';

const shortRentCars = [
  {
    brand: 'Toyota',
    model: 'Hiace',
    rentalPricePerDay: 2080.00,
    seat: 12,
    doors: 4,
    gearType: 'อัตโนมัติ',
    fuelType: 'เบนซิน',
    image: '/src/assets/images/cars/Hiace.png'
  },
  {
    brand: 'Hyundai',
    model: 'Staria',
    rentalPricePerDay: 3180.00,
    seat: 11,
    doors: 4,
    gearType: 'อัตโนมัติ',
    fuelType: 'เบนซิน',
    image: '/src/assets/images/cars/Staria.png'
  },
  {
    brand: 'Toyota',
    model: 'Sienta',
    rentalPricePerDay: 1320.00,
    seat: 7,
    doors: 4,
    gearType: 'อัตโนมัติ',
    fuelType: 'เบนซิน',
    image: '/src/assets/images/cars/Sienta.png'
  }
];
</script>

<style scoped>
@import '../assets/rent-service.css';

</style>