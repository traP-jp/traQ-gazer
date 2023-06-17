import { createApp, VueElement } from 'vue'
import './style.css'
import App from './App.vue'
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


const app = createApp(App)
app.use(router)
app.mount('#app')
