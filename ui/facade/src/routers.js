import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home.vue';
import Events from './components/Events.vue';

Vue.use(VueRouter);

export default new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: Home },
    { path: '/home', component: Home },
    { path: '/events', component: Events },
  ]
});

