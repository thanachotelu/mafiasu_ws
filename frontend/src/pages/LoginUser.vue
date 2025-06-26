<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { authService } from '../services/authService';
import NavBar1 from '../components/NavBar1.vue';

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
    
    // Check if user has user role instead of affiliate role
    
   console.log(response.roles);
 if (!response.roles?.includes('user')) {
    error.value = 'Access denied. User account required.';
    return;
}

    router.push('/car-lists');
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
  <NavBar1 />
  <div class="main-container">
    <!-- Add affiliator banner -->
    <div class="affiliator-banner">
      <p>Are you an affiliator? 
        <RouterLink to="/login-affiliator" class="affiliator-link">
          Login as Affiliator
        </RouterLink>
      </p>
    </div>
    
    <div class="login-page">
      <div class="login-container">
        <h1>Customer Login</h1>
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
          <RouterLink to="/user-form" class="nav-link">
            Register Now
          </RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import '../assets/user-login.css';

.affiliator-banner {
  background-color: #f8f9fa;
  padding: 10px 0;
  text-align: center;
  border-bottom: 1px solid #e9ecef;
  width: 100%;
}

.affiliator-banner p {
  margin: 0;
  font-size: 0.9rem;
  color: #6c757d;
}

.affiliator-link {
  color: #4CAF50;
  text-decoration: none;
  font-weight: 600;
  margin-left: 5px;
}

.affiliator-link:hover {
  text-decoration: underline;
}
</style>