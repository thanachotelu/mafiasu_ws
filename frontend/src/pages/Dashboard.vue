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
  import { ref, computed, onMounted } from 'vue'
  import { useRouter } from 'vue-router'
  import axios from 'axios'

  /* ------------------------------------------------------------------
  *  state
  * -----------------------------------------------------------------*/
  const router        = useRouter()
  const isLoading     = ref(false)
  const error         = ref('')

  // token ถูกเก็บไว้ตั้งแต่ตอน login ด้วย localStorage.setItem('token', tokenString)
  const storedToken   = ref(localStorage.getItem('token') || '')

  const user = ref({
    name       : localStorage.getItem('username') || '',  // ← ใช้ username ที่บันทึกไว้
    accessToken: storedToken.value
  });

  const carClickLogs  = ref([])
  const currentPage   = ref(1)
  const itemsPerPage  = 5

  /* ------------------------------------------------------------------
  *  computed
  * -----------------------------------------------------------------*/
  const totalPages = computed(() =>
    Math.ceil(carClickLogs.value.length / itemsPerPage)
  )

  const paginatedLogs = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage
    const end   = start + itemsPerPage
    return carClickLogs.value.slice(start, end)
  })

  /* ------------------------------------------------------------------
  *  methods
  * -----------------------------------------------------------------*/
  const fetchData = async () => {
    // ถ้าไม่มี token ให้เด้งกลับไปหน้า login (หรือหน้า home) เพื่อ re-auth
    if (!storedToken.value) {
      router.push({ name: 'Login' })
      return
    }

    isLoading.value = true
    try {
      /*--------- User data ---------*/
      const userRes = await axios.get('http://localhost:8000/api/v1/auth/user', {
        headers: { Authorization: `Bearer ${storedToken.value}` }
      })

      user.value = {
        name       : userRes.data.username,
        accessToken: storedToken.value
      }

      /*--------- Click-logs --------*/
      const logsRes = await axios.get('http://localhost:8000/api/v1/clicklogs', {
        headers: { Authorization: `Bearer ${storedToken.value}` }
      })

      carClickLogs.value = logsRes.data.map(log => ({
        carName    : `${log.brand} ${log.model}`,
        clickCount : log.click_count,
        latestClick: log.latest_click
      }))
    } catch (e) {
      error.value = 'ไม่สามารถโหลดข้อมูลได้ กรุณาลองใหม่'
      console.error(e)
      // กรณี token หมดอายุ/ไม่ถูกต้อง เคลียร์แล้วให้ไป login ใหม่
      if (e.response?.status === 401) {
        localStorage.removeItem('token')
        router.push({ name: 'Login' })
      }
    } finally {
      isLoading.value = false
    }
  }

  /* --------- pagination helpers ---------*/
  const prevPage = () => currentPage.value > 1              && currentPage.value--
  const nextPage = () => currentPage.value < totalPages.value && currentPage.value++

  /* --------- date helper ---------*/
  const formatDate = (iso) => {
    const opts = { year:'numeric', month:'short', day:'numeric', hour:'2-digit', minute:'2-digit' }
    return new Date(iso).toLocaleDateString('th-TH', opts)
  }

  /* --------- copy helper ---------*/
  const copyToken = async () => {
    try {
      await navigator.clipboard.writeText(user.value.accessToken)
    } catch { /* เงียบไว้หรือจะแจ้ง error ก็ได้ */ }
  }

  /* ------------------------------------------------------------------
  *  lifecycle
  * -----------------------------------------------------------------*/
  onMounted(fetchData)
  </script>
  
  <style scoped>
  @import '../assets/dashboard.css'
  </style>