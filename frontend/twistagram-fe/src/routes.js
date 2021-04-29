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
            path: '/home/:userId',
            name: 'userHome',
            component: () => import('./components/Home')
        },
        {
            path: '/:userId/profile',
            name: 'profile',
            component: () => import('./components/profile')
        },
        {
            path: '/home/:userId/search',
            name: 'search',
            component: () => import('./components/Search')
        }
        
    ]
});