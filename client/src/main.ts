import { createApp } from 'vue'
import './style.scss'
import 'vue-final-modal/style.css'
import App from './App.vue'
import router from './router'
import { createVfm } from 'vue-final-modal'

const app = createApp(App)
app.use(createVfm())
app.use(router)
app.mount('#app')
