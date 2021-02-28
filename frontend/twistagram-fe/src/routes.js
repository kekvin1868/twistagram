import Vue from 'vue'
import VueRouter from 'vue-router';

Vue.use(VueRouter);

export default new VueRouter({
    mode:'history',
    routes: [
        {
            path: '/',
            name: 'login',
            component: () => import('./components/login')
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('./components/register')
        },
        {
            path: '/:userId',
            name: 'userHome',
            component: () => import('./components/Home')
        }
        
    ]
});