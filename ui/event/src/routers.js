import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';
import Events from './components/Events';
import login from './components/Login';

Vue.use(VueRouter);

export default new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: Home },
    { path: '/home', component: Home },
    { path: '/events', component: Events },
    { path: '/login', component: login },
  ]
});

