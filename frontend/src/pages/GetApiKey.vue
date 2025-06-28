<template>
  <div class="form-container">
    <!-- Loading Popup -->
    <div v-if="showLoading" class="loading-overlay">
      <div class="loading-popup">
        <div class="loading-spinner"></div>
        <p class="loading-text">Please wait...</p>
      </div>
    </div>

    <div class="form-card">
      <!-- Form -->
      <div v-if="!showToken">
        <h2 class="form-title">Generate API Key</h2>
        
        <form @submit.prevent="handleSubmit" class="contact-form">
          <div class="form-group">
            <label for="name" class="form-label">Name</label>
            <input
              id="name"
              v-model="form.name"
              type="text"
              class="form-input"
              placeholder="Example: Joe"
              required
            />
          </div>

          <div class="form-group">
            <label for="email" class="form-label">Email</label> 
            <input
              id="email"
              v-model="form.email"
              type="email"
              class="form-input"
              placeholder="Example: joe@example.com"
              required
            />
          </div>

          <button type="submit" class="submit-btn" :disabled="isSubmitting">
            <span v-if="!isSubmitting">Submit</span>
            <span v-else>Submitting...</span>
          </button>
        </form>
      </div>

      <!-- Access Token Display -->
      <div v-else class="token-display">
        <div class="token-header">
          <h2 class="token-title">Your API Key is ready!</h2>
        </div>
        
        <div class="token-field">
          <input
            v-model="accessToken"
            type="text"
            class="token-input"
            readonly
          />
          <button @click="copyToken" class="copy-btn" :class="{ 'copied': tokenCopied }">
            {{ tokenCopied ? 'copied' : 'copy to clipboard' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: 'KeyForm',
  data() {
    return {
      form: {
        name: '',
        email: ''
      },
      isSubmitting: false,
      showLoading: false,
      showToken: false,
      accessToken: '',
      tokenCopied: false
    }
  },
  methods: {
    async handleSubmit() {
      this.isSubmitting = true;
      this.showLoading = true;
      try {
        // Wait at least 5 seconds before hiding the loading popup
        const loadingDelay = new Promise(resolve => setTimeout(resolve, 5000));
        // Start API request
        const apiRequest = axios.post("http://localhost:8000/api/v1/client/create", {
          name: this.form.name,
          email: this.form.email
        });

        // Wait for both the API and the delay to finish
        const [response] = await Promise.all([apiRequest, loadingDelay]);
        console.log('API response:', response.data);

        this.accessToken = response.data.api_key;
        this.showToken = true;
      } catch (error) {
        console.error('Error submitting form:', error);
        alert('Error submitting form. Please try again.');
      } finally {
        this.showLoading = false;
        this.isSubmitting = false;
      }
    },
    
    async copyToken() {
      try {
        await navigator.clipboard.writeText(this.accessToken);
        this.tokenCopied = true;
        setTimeout(() => {
          this.tokenCopied = false;
        }, 2000);
      } catch (error) {
        console.error('Failed to copy token:', error);
      }
    }
  }
}
</script>

<style scoped>
@import '../assets/getapikey.css';
</style>