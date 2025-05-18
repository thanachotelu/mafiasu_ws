<template>
  <nav class="navbar">
    <div class="logo">
      <img src="../assets/images/car-rental.png" alt="MafiaCar Logo">
      <h2>MafiaCar</h2>
    </div>
    <div class="nav-links">
      
      <RouterLink to="/" class="nav-link">Home</RouterLink>
      <RouterLink to="/car-lists" class="nav-link">Car Listing</RouterLink>
      <RouterLink to="/login-user" class="nav-link">Login/Register</RouterLink>

    </div>
  </nav>
</template>

<script setup>
</script>

<style scoped>
@import '../assets/navbar.css';
</style>