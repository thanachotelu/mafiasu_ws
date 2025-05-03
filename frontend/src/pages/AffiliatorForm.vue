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
.main-container {
  max-width: 100%; /* กำหนดความกว้างสูงสุดของหน้า */
  margin: 0; /* จัดให้อยู่กึ่งกลาง */
  padding: 10px; /* ระยะห่างจากขอบซ้ายขวา */
  width: 100%;
}


.form-container {
  width: 100%;
  /* min-height: 100vh; */
  height: fit-content;
  margin: 0;
  padding: 40px 0;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  
  /* เพิ่มส่วนนี้แทน background-color */
  background-image: url('../assets/images/affiliateForm.jpg');
  background-size: cover; /* ให้รูปเต็มพื้นที่ */
  background-position: center; /* จัดรูปให้อยู่กลาง */
  background-repeat: no-repeat; /* ไม่ให้รูปภาพซ้ำ */
  background-attachment: fixed; /* ให้พื้นหลังไม่เลื่อนตามหน้า */
}

.content-wrapper {
  width: 100%;
  max-width: 1200px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 30px;
}

.header-text {
  text-align: center;
  color: white;
}

h1 {
  margin: 0;
  font-size: 2.5rem;
  font-weight: 600;
  margin-bottom: 15px;
  text-shadow: 1px 1px 3px rgba(0, 0, 0, 0.2);
}

.welcome-text {
  margin: 0;
  font-size: 1.2rem;
  line-height: 1.6;
  opacity: 0.9;
}

.welcome-text a{
  color: #4CAF50;
  text-decoration: underline; 
  font-weight: 600; 
}

.white-box {
  width: 100%;
  height: 900px;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  display: flex;
  overflow: hidden;
}

.form-section {
  flex: 1; /* เปลี่ยนจาก 1 เป็น 2 */
  padding: 40px;
  overflow-y: auto;
}

.image-section {
  flex: 1; /* รักษาไว้ที่ 1 */
  display: flex;
  overflow: hidden;
}

.side-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
}

.placeholder-image {
  text-align: center;
  color: #666;
  padding: 20px;
}

form {
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
  color: #333;
}

input {
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
}

.name-fields, .address-fields {
  display: flex;
  gap: 10px;
}

.name-fields input, .address-fields input {
  flex: 1;
}

.required {
  color: red;
}

.required-note {
  font-size: 0.8rem;
  color: #666;
  margin: 0;
}

.register-btn {
  background-color: #4CAF50;
  color: white;
  padding: 14px;
  border: none;
  border-radius: 6px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  margin-top: 20px;
  transition: background-color 0.3s;
}

.register-btn:hover {
  background-color: #45a049;
}

@media (max-width: 768px) {
  .white-box {
    flex-direction: column;
    height: auto;
  }
  
  .image-section {
    flex: 1;
    display: flex;
    overflow: hidden; /* เพิ่มเพื่อไม่ให้รูปภาพล้นกรอบ */
  }

  .side-image {
    width: 100%;
    height: 100%;
    object-fit: cover; /* ให้รูปภาพเต็มพื้นที่โดยรักษาสัดส่วน */
    object-position: center; /* จัดตำแหน่งรูปให้อยู่กลาง */
  }
}
</style>