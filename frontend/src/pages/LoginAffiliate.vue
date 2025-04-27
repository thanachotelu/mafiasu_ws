<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()

const handleLogin = () => {
  // Demo credentials check
  if (username.value === 'Test' && password.value === 'Pass1234') {
    // Get NavBar component instance and update login state
    const navBar = document.querySelector('nav').__vueParentComponent.exposed
    navBar.isLoggedIn.value = true
    router.push('/dashboard') // Redirect to home page after successful login
  } else {
    errorMessage.value = 'Invalid username or password. Please try again.'
  }
}
</script>

<template>
  <div class="main-container">
    <div class="login-page">
      <div class="login-container">
        <h1>Affiliator Login</h1>
        <form @submit.prevent="handleLogin" class="login-form">
          <div class="form-group">
            <label for="username">Username / Email</label>
            <input
              type="text"
              id="username"
              v-model="username"
              required
            />
          </div>
          <div class="form-group">
            <label for="password">Password</label>
            <input
              type="password"
              id="password"
              v-model="password"
              required
            />
          </div>
          <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
          <button type="submit" class="login-button">Login</button>
        </form>
        <p class="register-link">
          don't have account yet? <RouterLink to="/affiliator-form" class="nav-link">register now</RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.main-container {
  max-width: 100%;
  margin: 0;
  padding: 20px;
  width: 100%;
}

.login-page {
  display: flex;
  justify-content: flex-start; /* Changed from flex-end to flex-start */
  align-items: center;
  height: 80vh;
  background-image: url('../assets/images/LoginAffiliator.png'); 
  background-size: cover;
  background-position: center;
  padding-left: 10%; /* Changed from padding-right to padding-left */
}

.login-container {
  width: 350px;
  padding: 40px;
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  margin-bottom: 30px;
  color: #333;
  font-size: 24px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

label {
  font-weight: 600;
  color: #555;
}

input {
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.3s;
}

input:focus {
  border-color: #4285f4;
  outline: none;
}

.login-button {
  padding: 12px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.3s;
  margin-top: 10px;
}

.login-button:hover {
  background-color: #338f36;
}

.register-link {
  text-align: center;
  margin-top: 20px;
  color: #666;
  font-size: 14px;
}

.register-link a {
  color: #4CAF50;
  text-decoration: none;
  font-weight: 600;
}

.register-link a:hover {
  text-decoration: underline;
}
.error-message {
  color: #ff4444;
  text-align: center;
  margin-top: 10px;
  font-size: 14px;
}
</style>