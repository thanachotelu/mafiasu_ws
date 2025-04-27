<script setup>
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const selectedTypes = ref([]);
const selectedBrands = ref([]);
const selectedYears = ref([]);
const minPrice = ref('');
const maxPrice = ref('');
const searchQuery = ref('');

const viewCarDetail = (carId) => {
  const selectedCar = cars.value.find(car => car.id === carId);
  router.push({
    name: 'CarDetail',
    params: { id: carId },
    query: { car: JSON.stringify(selectedCar) }
  });
};

const cars = ref([
  { id: 1, name: 'Toyota Corolla Cross', type: 'suv', brand: 'toyota', price: 859000, year: 2023, location: 'กรุงเทพมหานคร', km: 20000 },
  { id: 2, name: 'Honda City RS', type: 'sedan', brand: 'honda', price: 769000, year: 2022, location: 'ชลบุรี', km: 15000 },
  { id: 3, name: 'Ford Ranger Raptor', type: 'pickup', brand: 'ford', price: 1499000, year: 2023, location: 'ภูเก็ต', km: 8000 },
  { id: 4, name: 'MG ZS', type: 'suv', brand: 'mg', price: 659000, year: 2021, location: 'เชียงใหม่', km: 35000 },
  { id: 5, name: 'Mitsubishi Pajero Sport', type: 'suv', brand: 'mitsubishi', price: 1299000, year: 2020, location: 'นนทบุรี', km: 45000 },
  { id: 6, name: 'Isuzu D-Max', type: 'pickup', brand: 'isuzu', price: 789000, year: 2022, location: 'สมุทรปราการ', km: 25000 },
]);

const filterByYearRange = (carYear, selectedRange) => {
  const rangeMap = {
    '2023': [2023, 9999],
    '2020-2022': [2020, 2022],
    '2017-2019': [2017, 2019],
    '2014-2016': [2014, 2016],
  };
  return selectedRange.some(range => {
    const [start, end] = rangeMap[range];
    return carYear >= start && carYear <= end;
  });
};

const filteredCars = computed(() => {
  return cars.value.filter(car => {
    const matchType = selectedTypes.value.length ? selectedTypes.value.includes(car.type) : true;
    const matchBrand = selectedBrands.value.length ? selectedBrands.value.includes(car.brand) : true;
    const matchYear = selectedYears.value.length ? filterByYearRange(car.year, selectedYears.value) : true;
    const matchMinPrice = minPrice.value ? car.price >= parseInt(minPrice.value) : true;
    const matchMaxPrice = maxPrice.value ? car.price <= parseInt(maxPrice.value) : true;
    const matchSearch = searchQuery.value ? 
      car.name.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
      car.location.toLowerCase().includes(searchQuery.value.toLowerCase()) : true;
    
    return matchType && matchBrand && matchYear && matchMinPrice && matchMaxPrice && matchSearch;
  });
});
</script>

<template>
  <div class="main-container">
    <!-- Search Bar Section -->
    <div class="search-bar-container">
      <div class="search-bar">
        <input 
          type="text" 
          v-model="searchQuery" 
          placeholder="ค้นหารถยนต์ (ชื่อรถหรือสถานที่)" 
          class="search-input"
        />
        <button class="search-button">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
            <path d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"/>
          </svg>
        </button>
      </div>
    </div>

    <div class="search-container">
      <!-- Filter Section -->
      <div class="filter-section">
        <h2>ประเภทรถ</h2>
        <div class="filter-group">
          <label><input type="checkbox" value="eco" v-model="selectedTypes" /> รถ ECO-Car</label>
          <label><input type="checkbox" value="sedan" v-model="selectedTypes" /> รถเก๋ง (4ประตู Sedan)</label>
          <label><input type="checkbox" value="suv" v-model="selectedTypes" /> รถ SUV (7 ที่นั่ง)</label>
          <label><input type="checkbox" value="mpv" v-model="selectedTypes" /> รถอเนกประสงค์ (MPV)</label>
          <label><input type="checkbox" value="van" v-model="selectedTypes" /> รถตู้</label>
          <label><input type="checkbox" value="pickup" v-model="selectedTypes" /> รถกระบะ</label>
        </div>

        <h2>ยี่ห้อรถ</h2>
        <div class="filter-group">
          <label><input type="checkbox" value="toyota" v-model="selectedBrands" /> โตโยต้า (Toyota)</label>
          <label><input type="checkbox" value="honda" v-model="selectedBrands" /> ฮอนด้า (Honda)</label>
          <label><input type="checkbox" value="isuzu" v-model="selectedBrands" /> อีซูซุ (Isuzu)</label>
          <label><input type="checkbox" value="mitsubishi" v-model="selectedBrands" /> มิตซูบิชิ (Mitsubishi)</label>
          <label><input type="checkbox" value="ford" v-model="selectedBrands" /> ฟอร์ด (Ford)</label>
          <label><input type="checkbox" value="mg" v-model="selectedBrands" /> MG</label>
          <label><input type="checkbox" value="hyundai" v-model="selectedBrands" /> อุนได (Hyundai)</label>
        </div>

        <h2>ช่วงราคา</h2>
        <div class="price-range">
          <div class="price-inputs">
            <input type="number" v-model="minPrice" placeholder="ราคาต่ำสุด" />
            <span>ถึง</span>
            <input type="number" v-model="maxPrice" placeholder="ราคาสูงสุด" />
          </div>
        </div>

        <h2>ปีรถ</h2>
        <div class="filter-group">
          <label><input type="checkbox" value="2023" v-model="selectedYears" /> 2023-ปัจจุบัน</label>
          <label><input type="checkbox" value="2020-2022" v-model="selectedYears" /> 2020-2022</label>
          <label><input type="checkbox" value="2017-2019" v-model="selectedYears" /> 2017-2019</label>
          <label><input type="checkbox" value="2014-2016" v-model="selectedYears" /> 2014-2016</label>
        </div>
      </div>

      <!-- Car List Section -->
      <div class="car-list-section">
        <div class="header-section">
          <h1>รถทั้งหมด [{{ filteredCars.length }}]</h1>
        </div>
        <div class="car-grid">
          <div v-for="car in filteredCars" :key="car.id" class="car-card" @click="viewCarDetail(car.id)">
            <div class="car-image-placeholder">[รูป]</div>
            <div class="car-info">
              <h3>{{ car.name }}</h3>
              <p class="price">฿{{ car.price.toLocaleString() }}</p>
              <p class="details">{{ car.year }} · {{ car.km.toLocaleString() }} km · เกียร์อัตโนมัติ</p>
              <p class="location">{{ car.location }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.main-container {
  max-width: 100%;
  margin: 0;
  padding: 10px;
  width: 100%;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.search-bar-container {
  display: flex;
  justify-content: center;
  margin: 20px 0;
}

.search-bar {
  display: flex;
  max-width: 600px;
  width: 100%;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  border-radius: 30px;
  overflow: hidden;
}

.search-input {
  flex: 1;
  padding: 12px 20px;
  border: none;
  font-size: 1rem;
  outline: none;
}

.search-button {
  padding: 0 20px;
  border: none;
  background: #e53935;
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.search-button:hover {
  background: #c62828;
}

.search-container {
  display: flex;
  max-width: 1200px;
  margin: 20px auto;
  gap: 20px;
}

.filter-section {
  width: 270px;
  padding: 20px;
  background: #f8f8f8;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.filter-section h2 {
  font-size: 1.2rem;
  margin: 20px 0 10px 0;
  color: #333;
}

.filter-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.filter-group label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.price-range {
  margin: 15px 0;
}

.price-inputs {
  display: flex;
  align-items: center;
  gap: 10px;
}

.price-inputs input {
  width: 100px;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.car-list-section {
  flex: 1;
}

.header-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.header-section h1 {
  font-size: 1.5rem;
  color: #333;
}

.sort-options select {
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
}

.car-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
}

.car-card {
  background: white;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  transition: transform 0.2s;
}

.car-card:hover {
  transform: translateY(-5px);
}

.car-image-placeholder {
  height: 180px;
  background: #e0e0e0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
}

.car-info {
  padding: 15px;
}

.car-info h3 {
  margin: 0 0 10px 0;
  font-size: 1.1rem;
  color: #222;
}

.price {
  font-size: 1.2rem;
  font-weight: bold;
  color: #e53935;
  margin: 5px 0;
}

.details, .location {
  font-size: 0.9rem;
  color: #666;
  margin: 5px 0;
}

@media (max-width: 768px) {
  .search-container {
    flex-direction: column;
  }
  
  .filter-section {
    width: 100%;
  }
  
  .car-grid {
    grid-template-columns: 1fr;
  }

  .search-bar {
    max-width: 100%;
  }
}
</style>