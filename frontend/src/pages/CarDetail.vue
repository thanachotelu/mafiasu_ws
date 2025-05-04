<template>
  <div v-if="car.name">
      <div class="car-detail-container">
          <!-- Add return button -->
          <div class="return-button-container">
              <button @click="goBack" class="return-button">
                  <span class="arrow">←</span> ย้อนกลับ
              </button>
          </div>

          <div class="car-header">
              <h1 class="car-title">{{ car.name }}</h1>
              <div class="car-price">฿{{ car.price.toLocaleString() }}</div>
          </div>

          <div class="car-details">
            <div class="car-image">
              <img 
                :src="`/src/assets/images/cars/${carData.model}.png`" 
                :alt="car.name"
                              />
            </div>

            <div class="car-specs">
                <div class="spec-item">
                    <span class="spec-label">ประเภท</span>
                    <span class="spec-value">{{ car.type }}</span>
                </div>
                <div class="spec-item">
                    <span class="spec-label">เกียร์</span>
                    <span class="spec-value">{{ car.gear }}</span>
                </div>
                <div class="spec-item">
                    <span class="spec-label">ทะเบียน</span>
                    <span class="spec-value">{{ car.license_plate }}</span>
                </div>
                <div v-if="car.km" class="spec-item">
                    <span class="spec-label">ระยะทาง</span>
                    <span class="spec-value">{{ car.km.toLocaleString() }} km</span>
                </div>
                <div class="spec-item">
                    <span class="spec-label">ประเภทเชื้อเพลิง</span>
                    <span class="spec-value">{{ car.fuel }}</span>
                </div>
            </div>
          </div>

          <div class="car-description">
              <h2>รายละเอียดเพิ่มเติม</h2>
              <p>{{ car.name }} ทะเบียน {{ car.license_plate }} ราคาเช่า {{ car.price.toLocaleString() }} บาทต่อวัน</p>
          </div>
      </div>
  </div>
  <div v-else class="error-message">
      <h2>ไม่พบข้อมูลรถ</h2>
      <p>รหัสรถ: {{ id }}</p>
      <p>ข้อมูลรถ: {{ car }}</p>
      <p>Route Query: {{ route.query.car }}</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import axios from 'axios';

const router = useRouter();
const route = useRoute();

const car = ref({
  name: '',
  price: 0,
  gear:'',
  license_plate: '',
  brand: '',
  type: '',
  fuel: ''
});

const carData = ref({
  model: ''
});

const fetchCarData = async () => {
  try {
    const response = await axios.get(`http://localhost:8000/api/v1/cars/${route.params.id}`);
    const data = response.data;
    carData.value = data; // Store the full car data
    car.value = {
      name: `${data.brand} ${data.model}`,
      price: data.rental_price_per_day,
      gear: data.geartype,
      license_plate: data.license_plate,
      brand: data.brand.toLowerCase(),
      type: data.cartype,
      fuel: data.fueltype
    };
  } catch (error) {
    console.error('Failed to fetch car details:', error);
  }
};

onMounted(() => {
  fetchCarData();
});

const brandNames = {
  toyota: 'โตโยต้า',
  honda: 'ฮอนด้า',
  isuzu: 'อีซูซุ',
  mitsubishi: 'มิตซูบิชิ',
  ford: 'ฟอร์ด',
  hyundai: 'ฮุนได',
};

const goBack = () => {
    router.push('/car-lists');
};
</script>

<style scoped>
@import '../assets/cardetail.css';
</style>