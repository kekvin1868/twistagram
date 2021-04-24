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
            path: '/:userId/bookmark',
            name: 'bookmark',
            component: () => import('./components/bookmark')
        },
        {
            path: '/:userId/editProfile',
            name: 'editProfile',
            component: () => import('./components/editProfile')
        },
        {
            path: '/post/:postId',
            name: 'showPost',
            component: () => import('./components/ShowPost')
        }
        
    ]
});