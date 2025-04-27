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
              <div class="car-image-placeholder">
                  [รูปภาพรถ {{ car.name }}]
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

const fetchCarData = async () => {
  try {
    const response = await axios.get(`http://localhost:8000/api/v1/cars/${route.params.id}`);
    const carData = response.data;
    car.value = {
      name: `${carData.brand} ${carData.model}`,
      price: carData.rental_price_per_day,
      gear: carData.geartype,
      license_plate: carData.license_plate,
      brand: carData.brand.toLowerCase(),
      type: carData.cartype,
      fuel: carData.fueltype
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
  mg: 'เอ็มจี'
};

const getBrandName = (brandKey) => {
  return brandNames[brandKey] || brandKey;
};

const goBack = () => {
    router.push('/car-lists');
};
</script>

<style scoped>
.return-button-container {
    margin-bottom: 20px;
}

.return-button {
    display: flex;
    align-items: center;
    padding: 8px 16px;
    background: transparent;
    border: 1px solid #e53935;
    color: #e53935;
    border-radius: 20px;
    cursor: pointer;
    font-size: 0.9rem;
    transition: all 0.3s ease;
}

.return-button:hover {
    background: #e53935;
    color: white;
}

.arrow {
    margin-right: 8px;
    font-size: 1.2rem;
}

.car-detail-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.car-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.car-title {
  font-size: 1.8rem;
  color: #333;
  margin: 0;
}

.car-price {
  font-size: 1.5rem;
  font-weight: bold;
  color: #e53935;
}

.car-details {
  display: flex;
  gap: 30px;
  margin-bottom: 30px;
}

.car-image-placeholder {
  flex: 1;
  min-height: 400px;
  background: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  font-size: 1.2rem;
  color: #666;
}

.car-specs {
  flex: 0 0 300px;
  background: #f8f8f8;
  padding: 20px;
  border-radius: 8px;
}

.spec-item {
  display: flex;
  justify-content: space-between;
  padding: 12px 0;
  border-bottom: 1px solid #e0e0e0;
}

.spec-item:last-child {
  border-bottom: none;
}

.spec-label {
  color: #666;
}

.spec-value {
  font-weight: 500;
  color: #333;
}

.car-description {
  background: #f8f8f8;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.car-description h2 {
  font-size: 1.3rem;
  margin-top: 0;
  margin-bottom: 15px;
  color: #333;
}

.car-description p {
  margin: 0;
  color: #555;
  line-height: 1.6;
}

.contact-seller {
  text-align: center;
}

.contact-button {
  background: #e53935;
  color: white;
  border: none;
  padding: 12px 30px;
  font-size: 1rem;
  border-radius: 30px;
  cursor: pointer;
  transition: background 0.3s;
}

.contact-button:hover {
  background: #c62828;
}

@media (max-width: 768px) {
  .car-details {
      flex-direction: column;
  }

  .car-specs {
      flex: 1;
  }

  .car-image-placeholder {
      min-height: 250px;
  }
}
</style>