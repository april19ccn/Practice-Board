import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import { Button, Picker } from 'vant';
import 'vant/lib/index.css';

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.use(Button);
app.use(Picker);

app.mount('#app')
