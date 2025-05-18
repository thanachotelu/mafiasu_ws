<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { authService } from '../services/authService';
import NavBar1 from '../components/NavBar1.vue';

const router = useRouter();
const isLoading = ref(false);
const error = ref('');
const showPassword = ref(false); // Add this line

const formData = ref({
  firstname: '',
  lastname: '',
  username: '',
  phonenumber: '',
  email: '',
  password: '',
  role: 'Affiliator'
});

const handleLogin = async (e) => {
  e.preventDefault();
  isLoading.value = true;
  error.value = '';

  try {
    await authService.register(formData.value);
    router.push('/login-affiliator');
  } catch (err) {
    error.value = err.message;
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <NavBar1 />
  <div class="main-container">
    <div class="form-container">
      <div class="content-wrapper">
        <div class="header-text">
          <h1>Affiliate Registration Form</h1>
          <p class="welcome-text">Welcome to affiliate program of MafiaCar. You can start by<br>
          registering your information below.
          <br>
          already have account? <RouterLink to="/login-affiliator" class="nav-link">Login now</RouterLink>
          </p>
        </div>
        <div class="white-box">
          <div class="form-section">
            <form @submit="handleLogin">
              <div v-if="error" class="error-message">
                {{ error }}
              </div>
              <div class="form-group">
                <label>Name</label>
                <div class="name-fields">
                  <input 
                    type="text" 
                    v-model="formData.firstname"
                    placeholder="First Name"
                    required
                  >
                  <input 
                    type="text" 
                    v-model="formData.lastname"
                    placeholder="Last Name"
                    required
                  >
                </div>
              </div>

              <div class="form-group">
                <label>Username <span class="required">*</span></label>
                <input 
                  type="text"
                  v-model="formData.username"
                  required
                >
                <p class="required-note">* Required</p>
              </div>

              <div class="form-group">
                <label>Phone number</label>
                <input 
                  type="tel"
                  v-model="formData.phonenumber"
                  required
                >
              </div>

              <div class="form-group">
                <label>Account Email <span class="required">*</span></label>
                <input 
                  type="email"
                  v-model="formData.email"
                  required
                >
                <p class="required-note">* Required</p>
              </div>

              <div class="form-group">
                <label>Password <span class="required">*</span></label>
                <div class="password-field">
                  <input 
                    :type="showPassword ? 'text' : 'password'"
                    v-model="formData.password"
                    required
                  >
                  <button 
                    type="button" 
                    class="toggle-password"
                    @click="showPassword = !showPassword"
                  >
                    {{ showPassword ? '🔓' : '🔒' }}
                  </button>
                </div>
                <p class="required-note">* Required</p>
              </div>

              <button 
                type="submit" 
                class="register-btn"
                :disabled="isLoading"
              >
                {{ isLoading ? 'Registering...' : 'Register now' }}
              </button>
            </form>
          </div>
          <div class="image-section">
            <img 
              src="../assets/images/affiliateForm2.jpg" 
              alt="Affiliate Program"
              class="side-image"
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import '../assets/affiliate-form.css';
</style>