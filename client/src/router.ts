import Home from "./pages/Home.vue"
import WordAdd from './pages/WordAdd.vue'
import WordList from './pages/WordList.vue'

import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/words",
    name: "WordList",
    component: WordList
  },
  {
    path: "/words/add",
    name: "WordAdd",
    component: WordAdd
  },
  
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router