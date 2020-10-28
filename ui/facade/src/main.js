import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './routers'
import axios from 'axios'

Vue.config.productionTip = false;
axios.defaults.baseURL = 'http://localhost:8090'

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
