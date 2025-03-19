import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Login from '../views/Login.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/dashboard'
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: Dashboard,
      meta: { requiresAuth: true }
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/products/:id',
      name: 'ProductDetail',
      component: () => import('../views/ProductDetail.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/products/:id/edit',
      name: 'EditProduct',
      component: () => import('../views/EditProduct.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

export default router