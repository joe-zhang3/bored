import Vue from 'vue'
import App from './App.vue'
import router from './routers'
import axios from 'axios'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

// Then import Bootstrap and BootstrapVue SCSS files (order is important)
//import "./app.scss"
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)
Vue.use(IconsPlugin)

Vue.config.productionTip = false;
axios.defaults.baseURL = 'http://event_app:8090'

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

