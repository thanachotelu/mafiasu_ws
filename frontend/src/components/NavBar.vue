<template>
  <nav class="navbar">
    <div class="logo">
      <img src="../assets/images/car-rental.png" alt="MafiaCar Logo">
      <h2>MafiaCar</h2>
    </div>
    <div class="nav-links">
      <template v-if="!isLoggedIn">
        <RouterLink to="/" class="nav-link">Home</RouterLink>
      </template>

      <RouterLink to="/car-lists" class="nav-link">Car Listing</RouterLink>
      
      <template v-if="!isLoggedIn">
        <RouterLink to="/login-affiliator" class="nav-link">Login As Affiliator</RouterLink>
      </template>
      
      <template v-else>
        <RouterLink to="/api-lists" class="nav-link">APIs Documentation</RouterLink>
        <RouterLink to="/dashboard" class="nav-link">Dashboard</RouterLink>
        <a href="#" @click.prevent="handleLogout" class="nav-link">Logout</a>
      </template>
    </div>
    <!-- Add popup message -->
    <div v-if="showPopup" class="popup" :class="{ 'fade-out': isClosing }">
      Successfully logged out!<br>ออกจากระบบเรียบร้อยแล้ว!
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'  // Add this import

const router = useRouter()  // Add this line
const isLoggedIn = ref(false)
const showPopup = ref(false)
const isClosing = ref(false)

const handleLogout = () => {
  isLoggedIn.value = false
  showPopup.value = true
  isClosing.value = false

  // Start fade out animation after 5 seconds
  setTimeout(() => {
    isClosing.value = true
  }, 5000)

  // Hide popup and redirect after fade out animation
  setTimeout(() => {
    router.push('/')  // Redirect to home page
  }, 500)
  setTimeout(() => {
    showPopup.value = false
  }, 2500)
}

defineExpose({ isLoggedIn })
</script>

<style scoped>
@import '../assets/navbar.css';
</style>