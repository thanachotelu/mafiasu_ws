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
  
  const router = useRouter();
  
  // ข้อมูลผู้ใช้
  const user = ref({
    name: 'John Doe',
    accessToken: 'Aff_containereyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE2MTYyMzkwMjJ9.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c'
  });
  
  // ข้อมูล Click Logs ตัวอย่าง
  const carClickLogs = ref([
    { carName: 'Toyota Corolla Cross', clickCount: 120, latestClick: '2025-04-25T15:30:00' },
    { carName: 'Honda City RS', clickCount: 75, latestClick: '2025-04-24T10:45:00' },
    { carName: 'Ford Ranger Raptor', clickCount: 90, latestClick: '2025-04-23T17:20:00' },
    { carName: 'MG ZS', clickCount: 45, latestClick: '2025-04-22T11:10:00' },
    { carName: 'Mitsubishi Pajero Sport', clickCount: 60, latestClick: '2025-04-21T09:00:00' },
    { carName: 'Isuzu MU-X', clickCount: 88, latestClick: '2025-04-20T14:15:00' },
    { carName: 'Nissan Almera', clickCount: 30, latestClick: '2025-04-19T08:50:00' },
    { carName: 'Mazda CX-5', clickCount: 100, latestClick: '2025-04-18T16:05:00' },
  ]);
  
  const currentPage = ref(1);
  const itemsPerPage = 5;
  
  const totalPages = computed(() => {
    return Math.ceil(carClickLogs.value.length / itemsPerPage);
  });
  
  const paginatedLogs = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage;
    const end = start + itemsPerPage;
    return carClickLogs.value.slice(start, end);
  });
  
  const prevPage = () => {
    if (currentPage.value > 1) currentPage.value--;
  };
  
  const nextPage = () => {
    if (currentPage.value < totalPages.value) currentPage.value++;
  };
  
  const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' };
    return new Date(dateString).toLocaleDateString('th-TH', options);
  };
  
  onMounted(() => {
    // โหลดข้อมูลจริงได้ที่นี่ในอนาคต
  });
  </script>
  
  <style scoped>
  .dashboard-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    background-color: #f5f5f5;
  }
  
  .dashboard-header {
    margin-bottom: 30px;
  }
  
  .dashboard-header h1 {
    color: #333;
    margin-bottom: 20px;
  }
  
  .access-token-section {
    background: #ffffff;
    padding: 20px;
    border-radius: 8px;
    margin-bottom: 30px;
  }
  
  .token-display {
    display: flex;
    align-items: center;
    gap: 10px;
    margin: 10px 0;
  }
  
  .token-display code {
    flex: 1;
    padding: 10px;
    background: #fff;
    border: 1px solid #ddd;
    border-radius: 4px;
    white-space: pre-wrap;
    /* <<< เพิ่มบรรทัดนี้ */
    word-break: break-word;
    /* <<< เพิ่มบรรทัดนี้ */
  }
  
  .copy-btn {
    padding: 10px 15px;
    background: #d21d1d;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background 0.3s;
  }
  
  .copy-btn:hover {
    background: #000000;
  }
  
  .token-note {
    color: #666;
    font-size: 0.9rem;
  }
  
  .click-logs-section {
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    padding: 20px;
  }
  
  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
  
  .filter-controls {
    display: flex;
    gap: 10px;
  }
  
  .filter-controls select,
  .search-input {
    padding: 8px 12px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }
  
  .logs-table-container {
    overflow-x: auto;
  }
  
  .logs-table {
    width: 100%;
    border-collapse: collapse;
  }
  
  .logs-table th,
  .logs-table td {
    padding: 12px 15px;
    text-align: left;
    border-bottom: 1px solid #eee;
  }
  
  .logs-table th {
    background: #f8f9fa;
    font-weight: 600;
  }
  
  .status-badge {
    padding: 5px 10px;
    border-radius: 20px;
    font-size: 0.8rem;
    font-weight: 500;
  }
  
  .status-badge.completed {
    background: #d4edda;
    color: #155724;
  }
  
  .status-badge.pending {
    background: #fff3cd;
    color: #856404;
  }
  
  .status-badge.failed {
    background: #f8d7da;
    color: #721c24;
  }
  
  .no-logs {
    text-align: center;
    padding: 30px;
    color: #666;
  }
  
  .pagination {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 15px;
    margin-top: 20px;
  }
  
  .page-btn {
    padding: 8px 16px;
    background: #f8f9fa;
    border: 1px solid #ddd;
    border-radius: 4px;
    cursor: pointer;
  }
  
  .page-btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
  
  @media (max-width: 768px) {
    .section-header {
      flex-direction: column;
      align-items: flex-start;
      gap: 10px;
    }
  
    .filter-controls {
      width: 100%;
    }
  
    .logs-table th,
    .logs-table td {
      padding: 8px;
      font-size: 0.9rem;
    }
  }
  
  .access-token-section i.fa-key {
    color: gold;
  }
  </style>