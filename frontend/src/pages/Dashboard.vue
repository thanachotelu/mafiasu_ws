<template>
    <div class="dashboard-container">
      <!-- ส่วนหัวหน้า Dashboard -->
      <div class="dashboard-header">
        <h1>👋 ยินดีต้อนรับ, {{ user.name }}</h1>
        <div class="access-token-section">
          <h2> 🔑 Access Token</h2>
          <div class="token-display">
            <code>{{ user.accessToken }}</code>
            <button @click="copyToken" class="copy-btn">
              <i class="fas fa-copy"></i> คัดลอก
            </button>
          </div>
        </div>
      </div>
  
      <!-- ส่วนแสดง Click Logs -->
      <div class="click-logs-section">
        <div class="section-header">
          <h2>🖱️สรุปรายการ Click Logs</h2>
        </div>
  
        <div class="logs-table-container">
          <table class="logs-table">
            <thead>
              <tr>
                <th>ชื่อรถ</th>
                <th>จำนวนการคลิก</th>
                <th>การคลิกล่าสุด</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(log, index) in paginatedLogs" :key="index">
                <td>{{ log.carName }}</td>
                <td>{{ log.clickCount }}</td>
                <td>{{ formatDate(log.latestClick) }}</td>
              </tr>
            </tbody>
          </table>
  
          <div v-if="paginatedLogs.length === 0" class="no-logs">
            <p>ไม่พบข้อมูล Click Logs</p>
          </div>
        </div>
  
        <div class="pagination" v-if="paginatedLogs.length > 0">
          <button @click="prevPage" :disabled="currentPage === 1" class="page-btn">
            ก่อนหน้า
          </button>
          <span>หน้า {{ currentPage }} จาก {{ totalPages }}</span>
          <button @click="nextPage" :disabled="currentPage === totalPages" class="page-btn">
            ถัดไป
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, computed, onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import axios from 'axios';
  
  const router = useRouter();
  const isLoading = ref(false);
  const error = ref('');
  
  // ข้อมูลผู้ใช้
  const user = ref({
    name: '',
    accessToken: ''
  });
  
  // ข้อมูล Click Logs
  const carClickLogs = ref([]);
  const currentPage = ref(1);
  const itemsPerPage = 5;
  
  // Pagination computeds
  const totalPages = computed(() => {
    return Math.ceil(carClickLogs.value.length / itemsPerPage);
  });
  
  const paginatedLogs = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage;
    const end = start + itemsPerPage;
    return carClickLogs.value.slice(start, end);
  });
  
  // ดึงข้อมูลผู้ใช้และ Click Logs
  const fetchData = async () => {
    isLoading.value = true;
    try {
      // Get user data from token
      const token = localStorage.getItem('token');
      const userResponse = await axios.get('http://localhost:8000/api/v1/auth/user', {
        headers: { Authorization: `Bearer ${token}` }
      });
      user.value = {
        name: userResponse.data.username,
        accessToken: token
      };
  
      // Get click logs
      const logsResponse = await axios.get('http://localhost:8000/api/v1/clicklogs', {
        headers: { Authorization: `Bearer ${token}` }
      });
      carClickLogs.value = logsResponse.data.map(log => ({
        carName: `${log.brand} ${log.model}`,
        clickCount: log.click_count,
        latestClick: log.latest_click
      }));
    } catch (err) {
      error.value = 'Failed to load data. Please try again.';
      console.error('Error fetching data:', err);
    } finally {
      isLoading.value = false;
    }
  };
  
  // Pagination methods
  const prevPage = () => {
    if (currentPage.value > 1) currentPage.value--;
  };
  
  const nextPage = () => {
    if (currentPage.value < totalPages.value) currentPage.value++;
  };
  
  // Format date helper
  const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
    return new Date(dateString).toLocaleDateString('th-TH', options);
  };
  
  // Copy token function
  const copyToken = () => {
    navigator.clipboard.writeText(user.value.accessToken);
  };
  
  onMounted(() => {
    fetchData();
  });
  </script>
  
  <style scoped>
  @import '../assets/dashboard.css'
  </style>