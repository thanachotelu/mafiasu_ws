<template>
  <nav class="navbar">
    <div class="logo">
      <img src="../assets/images/car-rental.png" alt="MafiaCar Logo">
      <h2>MafiaCar</h2>
    </div>
    <div class="nav-links">
      
      <RouterLink to="/" class="nav-link">Home</RouterLink>
      <RouterLink to="/car-lists" class="nav-link">Car Listing</RouterLink>
      <RouterLink to="/get-api-key" class="nav-link">Get API Key</RouterLink>

      <!-- Links when NOT logged in -->
      <template v-if="!isAuthenticated">
        <RouterLink to="/login-user" class="nav-link">Login/Register</RouterLink>
      </template>

      <!-- Links when logged in -->
      <template v-else>
        <RouterLink to="/api-lists" class="nav-link">APIs Documentation</RouterLink>
        <!-- Show Dashboard only for Affiliator role -->
        <RouterLink v-if="userRole === 'Affiliator'" to="/dashboard" class="nav-link">Dashboard</RouterLink>
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
import { useRouter } from 'vue-router'
import { authService, isAuthenticated, userRole } from '../services/authService' // Import userRole

const router = useRouter()
const showPopup = ref(false)
const isClosing = ref(false)

const handleLogout = () => {
  // Use the authService logout method (this will automatically update isAuthenticated)
  authService.logout()
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
</script>

<style scoped>
@import '../assets/navbar.css';
</style>