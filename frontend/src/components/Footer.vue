<template>
    <footer class="footer">
      <div class="footer-content">
        <div class="footer-left">
          <h2 class="footer-title">MafiaCar</h2>
          <p class="footer-description">
            ร้านนำรถ MafiaCar<br />
            มอบบริการรถเช่าคุณภาพ ด้วยความมั่นใจและไว้วางใจ<br />
            พร้อมให้บริการทุกการเดินทางของคุณ
          </p>
        </div>
        <div class="footer-right">
          <h3 class="contact-title">ติดต่อเรา</h3>
          <ul class="contact-list">
            <li class="contact-item">
              <img src="../assets/images/phone-call.png" alt="Phone" class="contact-icon" />
              012-345-6789
            </li>
            <li class="contact-item">
              <img src="../assets/images/email.png" alt="Email" class="contact-icon" />
              Mafiacar_su@gmail.com
            </li>
            <li class="contact-item">
              <img src="../assets/images/gps.png" alt="Address" class="contact-icon" />
              6 ถนนราชดำเนิน ตำบลพระปฐมเจดีย์ อำเภอเมืองนครปฐม นครปฐม 73000
            </li>
          </ul>
          <a href="#" class="button footer-button">ดูเพิ่มเติม</a>
        </div>
      </div>
    </footer>
</template>

<script setup>
// Component logic
</script>

<style scoped>
@import '../assets/footer.css';

.contact-icon {
  width: 16px;
  height: 16px;
  margin-right: 8px;
  vertical-align: middle;
}

.contact-item {
  display: flex;
  align-items: center;
}
</style>