<template>
    <div>
        <MainBanner />
        <div class="services-section">
            <MainServices />
        </div>
        <div class="affiliate-banner">
            <div class="affiliate-content">
                <h2 class="affiliate-title">ร่วมเป็นพาร์ทเนอร์กับเรา</h2>
                <p class="affiliate-text">เข้าร่วมโปรแกรมพาร์ทเนอร์กับ MafiaCar และสร้างรายได้ไปด้วยกัน</p>
                <div class="affiliate-buttons">
                    <RouterLink to="/affiliator-form" class="affiliate-button register">สมัครเป็นพาร์ทเนอร์</RouterLink>
                    <RouterLink to="/login-affiliator" class="affiliate-button login">เข้าสู่ระบบพาร์ทเนอร์</RouterLink>
                </div>
            </div>
            <div class="affiliate-image"></div>
        </div>
        <Footer />
    </div>
</template>

<script setup>
import MainBanner from '../components/MainBanner.vue';
import MainServices from '../components/MainServices.vue';
import Footer from '../components/Footer.vue';
</script>

<style scoped>
@import '../assets/affiliate-banner.css';

.services-section {
    margin-top: 40px;
}
</style>