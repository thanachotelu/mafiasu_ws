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
          <h2>🖱️ สรุปรายการ Click Logs</h2>
        </div>
  
        <div class="logs-table-container">
          <table class="logs-table">
            <thead>
              <tr>
                <th>Endpoint</th>
                <th>Method</th>
                <th>Timestamp</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(log, index) in paginatedLogs" :key="index">
                <td>{{ log.Endpoint }}</td>
                <td>{{ log.Method }}</td>
                <td>{{ formatDate(log.Timestamp) }}</td>
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
  import { ref, computed, onMounted } from 'vue'
  import { useRouter } from 'vue-router'
  import axios from 'axios'

  const router        = useRouter()
  const isLoading     = ref(false)
  const error         = ref('')

  const storedToken   = ref(localStorage.getItem('token') || '')

  const user = ref({
    name       : localStorage.getItem('username') || '',
    accessToken: storedToken.value
  });

  // เปลี่ยนชื่อ array และโครงสร้าง log
  const clickLogs     = ref([])
  const currentPage   = ref(1)
  const itemsPerPage  = 5

  const totalPages = computed(() =>
    Math.ceil(clickLogs.value.length / itemsPerPage)
  )

  const paginatedLogs = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage
    const end   = start + itemsPerPage
    return clickLogs.value.slice(start, end)
  })

  const fetchData = async () => {
    if (!storedToken.value) {
      router.push({ name: 'Login' })
      return
    }

    isLoading.value = true
    try {
      // ดึง user_id จาก localStorage หรือ context
      const userId = localStorage.getItem('user_id')
      const logsRes = await axios.get(`http://localhost:8000/api/v1/client/getlogs/${userId}`, {
        headers: { Authorization: `Bearer ${storedToken.value}` }
      })
      clickLogs.value = logsRes.data
    } catch (e) {
      error.value = 'ไม่สามารถโหลดข้อมูลได้ กรุณาลองใหม่'
      if (e.response?.status === 401) {
        localStorage.removeItem('token')
        router.push({ name: 'Login' })
      }
    } finally {
      isLoading.value = false
    }
  }

  const prevPage = () => currentPage.value > 1              && currentPage.value--
  const nextPage = () => currentPage.value < totalPages.value && currentPage.value++

  const formatDate = (iso) => {
    const opts = { year:'numeric', month:'short', day:'numeric', hour:'2-digit', minute:'2-digit' }
    return new Date(iso).toLocaleDateString('th-TH', opts)
  }

  const copyToken = async () => {
    try {
      await navigator.clipboard.writeText(user.value.accessToken)
    } catch {}
  }

  onMounted(fetchData)
  </script>
  
  <style scoped>
  @import '../assets/dashboard.css'
  </style>