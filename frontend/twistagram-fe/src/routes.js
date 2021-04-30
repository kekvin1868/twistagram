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
            path: '/:userId/profile/:vistId',
            name: 'profile',
            component: () => import('./components/profile')
        },
        {
<<<<<<< HEAD
            path: '/home/:userId/search',
=======
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
            path: '/post/:postId/:userId',
            name: 'showPost',
            component: () => import('./components/ShowPost')
        },
        {
            path: '/updatePost/:postId',
            name: 'editPost',
            component: () => import('./components/UpdatePost')
        },
        {
            path: '/home/:userID/search',
>>>>>>> eaa3ab2cf3f706741d59d8a073f38ba56aff3af7
            name: 'search',
            component: () => import('./components/Search')
        }
        
    ]
});