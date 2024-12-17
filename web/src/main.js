import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import { router } from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import { install } from '~/init';

import "normalize.css";
import '~/assets/css/reset.css'
import '~/assets/css/common.css'
import '~/assets/css/tailwind.css'

import 'nprogress/nprogress.css'


const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

app.use(router)
app.use(ElementPlus)

install(app)

// 确保在初始化 Pinia 和其他插件之后再导入权限控制
import "./permission"

app.mount('#app')
