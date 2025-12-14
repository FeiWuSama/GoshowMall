import { createApp } from 'vue'
import { createPinia } from 'pinia'
import AntDesign from 'ant-design-vue'
import 'ant-design-vue/dist/reset.css'

import App from './App.vue'
import router from './router'
import { useAuthStore } from './stores/auth'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(AntDesign)



app.mount('#app')
