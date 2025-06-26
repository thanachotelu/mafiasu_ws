import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'
import "./services/axios-interceptor";

createApp(App).use(router).mount('#app')