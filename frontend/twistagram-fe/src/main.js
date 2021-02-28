import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import VueRouter from 'vue-router'
import Routes from './routes.js';

Vue.use(VueRouter);

import axios from 'axios';
window.axios = axios;

new Vue({
  vuetify,
  router: Routes,
  render: h => h(App)
}).$mount('#app')
