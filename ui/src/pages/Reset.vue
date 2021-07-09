import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '../pages/Home.vue'
import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
 
const routes: Array<RouteRecordRaw> = [
  { path: '/', component: Home},
  { path: '/login', component: Login},
  { path: '/register', component: Register},
]
 
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})
 
export default router
