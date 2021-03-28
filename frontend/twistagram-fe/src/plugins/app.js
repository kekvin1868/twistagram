const Vue = require('vue');
import VueRouter from 'vue-router';
Vue.use(VueRouter);

import login from './components/home.vue'
import register from './components/register.vue';

const router = new VueRouter({
    routes: [
        {path : '/', component: login},
        {path : '/registerPage', component: register}
    ]
});

new Vue({
    router
}).$mount('#wrapper')