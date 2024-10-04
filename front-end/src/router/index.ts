// src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router';
import { ref } from 'vue';
import Login from '@/views/LoginForm.vue';
// import Register from '@/views/Register.vue';
// import Dashboard from '@/views/Dashboard.vue';

const isLoggedIn = ref(false); // Replace with your actual authentication logic

const routes = [
  { path: '/', name: 'Login', component: Login },
  { path: '/login', name: 'Login', component: Login },
  // { path: '/dashboard', name: 'Dashboard', component: Dashboard },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Navigation guard to check authentication state
router.beforeEach((to, from, next) => {
  if (to.path !== '/login' && !isLoggedIn.value) {
    next('/login'); // Redirect to login if not authenticated
  } else {
    next(); // Allow navigation
  }
});
export default router;
