import { createRouter, createWebHistory } from 'vue-router';
import Homepage from '../pages/Homepage.vue';
import LoginAffiliate from '../pages/LoginAffiliate.vue';
import AffiliatorForm from '../pages/AffiliatorForm.vue';
import CarLists from '../pages/CarLists.vue';
import Dashboard from '../pages/Dashboard.vue';
import CarDetail from '../pages/CarDetail.vue';

import ShortRent from '../pages/ShortRent.vue';
import LongRent from '../pages/LongRent.vue';
import CorporateRent from '../pages/CorporateRent.vue';
import ChauffeurRent from '../pages/ChauffeurRent.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Homepage, // Use Homepage.vue for the root route
  },
  {
    path: '/car-detail/:id',
    name: 'CarDetail',
    component: CarDetail,
    props: (route) => ({
      id: route.params.id,
      carData: route.query.car ? JSON.parse(route.query.car) : null
    })
  },
  {
    path: '/short-rent-service',
    name: 'ShortRent',
    component: ShortRent,
  },
  {
    path: '/long-rent-service',
    name: 'LongRent',
    component: LongRent,
  },
  {
    path: '/corporate-rent-service',
    name: 'CorporateRent',
    component: CorporateRent,
  },
  {
    path: '/chauffeur-service',
    name: 'ChauffeurRent',
    component: ChauffeurRent,
  },
  {
    path: '/login-affiliator',
    name: 'LoginAffiliate',
    component: LoginAffiliate,
  },
  {
    path: '/affiliator-form',
    name: 'AffiliatorForm',
    component: AffiliatorForm,
  },
  {
    path: '/car-lists',
    name: 'CarLists',
    component: CarLists,
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;