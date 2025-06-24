<template>
    <div>
        <MainBanner />
        <div class="services-section">
            <MainServices />
        </div>
        <AffiliateBanner />
        <Footer />
    </div>
</template>

<script setup>
import MainBanner from '../components/MainBanner.vue';
import MainServices from '../components/MainServices.vue';
import AffiliateBanner from '../components/AffiliateBanner.vue';
import Footer from '../components/Footer.vue';
</script>

<style scoped>
.services-section {
    margin-top: 40px;
}
</style>