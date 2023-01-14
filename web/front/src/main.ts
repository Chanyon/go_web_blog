import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './style.css'
import App from './App.vue'

import FloatingVue from 'floating-vue'
import 'floating-vue/dist/style.css'

import router from './router'
const pinia = createPinia();
const app = createApp(App);

app.use(pinia);
app.use(FloatingVue);
app.use(router);
app.mount('#app');
