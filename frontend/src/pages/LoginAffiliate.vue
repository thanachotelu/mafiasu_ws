<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { authService } from '../services/authService';

const router = useRouter();
const isLoading = ref(false);
const error = ref('');
const showPassword = ref(false);

const formData = ref({
  username: '',
  password: ''
});

const handleLogin = async (e) => {
  e.preventDefault();
  
  // Form validation
  if (!formData.value.username || !formData.value.password) {
    error.value = 'Please fill in all fields';
    return;
  }

  isLoading.value = true;
  error.value = '';

  try {
    const response = await authService.login(formData.value);
    
    // Check if user has affiliate role
    const user = response.user;
    if (!user.roles?.includes('Affiliator')) {
      error.value = 'Access denied. Affiliate account required.';
      return;
    }

    router.push('/dashboard');
  } catch (err) {
    error.value = err.message;
  } finally {
    isLoading.value = false;
  }
};

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value;
};
</script>

<template>
  <div class="main-container">
    <div class="login-page">
      <div class="login-container">
        <h1>Affiliator Login</h1>
        <form @submit="handleLogin" class="login-form">
          <div v-if="error" class="error-message" role="alert">
            {{ error }}
          </div>

          <div class="form-group">
            <label for="username">Username</label>
            <input
              type="text"
              id="username"
              v-model="formData.username"
              required
              autocomplete="username"
              :disabled="isLoading"
            />
          </div>

          <div class="form-group">
            <label for="password">Password</label>
            <div class="password-input">
              <input
                :type="showPassword ? 'text' : 'password'"
                id="password"
                v-model="formData.password"
                required
                autocomplete="current-password"
                :disabled="isLoading"
              />
              <button 
                type="button"
                class="toggle-password"
                @click="togglePasswordVisibility"
                :aria-label="showPassword ? 'Hide password' : 'Show password'"
              >
                {{ showPassword ? '🔓' : '🔒' }}
              </button>
            </div>
          </div>

          <button 
            type="submit" 
            class="login-button"
            :disabled="isLoading"
          >
            {{ isLoading ? 'Logging in...' : 'Login' }}
          </button>
        </form>

        <p class="register-link">
          Don't have an account yet? 
          <RouterLink to="/affiliator-form" class="nav-link">
            Register now
          </RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import '../assets/affiliate-login.css';
</style>