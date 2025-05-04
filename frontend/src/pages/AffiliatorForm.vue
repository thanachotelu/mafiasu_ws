<script setup>

</script>

<template>
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
            <form>
              <div class="form-group">
                <label>Name</label>
                <div class="name-fields">
                  <input type="text" placeholder="First Name">
                  <input type="text" placeholder="Last Name">
                </div>
              </div>

              <div class="form-group">
                <label>Username</label>
                <input type="text">
              </div>

              <div class="form-group">
                <label>Phone number</label>
                <input type="tel">
              </div>

              <div class="form-group">
                <label>Account Email <span class="required">*</span></label>
                <input type="email">
                <p class="required-note">* Required</p>
              </div>

              <div class="form-group">
                <label>Password <span class="required">*</span></label>
                <input type="password">
                <p class="required-note">* Required</p>
              </div>

              <button type="submit" class="register-btn">Register now</button>
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