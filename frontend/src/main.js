import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/reset.css'
app.use(Antd)

app.config.warnHandler = function(msg, vm, trace) {
    
}

app.mount('#app')