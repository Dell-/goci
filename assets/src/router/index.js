import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

// route-level code splitting
// const createListView = id => () => import('../views/CreateListView').then(m => m.default(id));
// const ProjectList = () => import('../views/ProjectList.vue');
const LoginView = () => import('../views/LoginView.vue');

// const UserView = () => import('../views/UserView.vue');

export function createRouter () {
  return new Router({
    mode: 'history',
    fallback: false,
    scrollBehavior: () => ({ y: 0 }),
    routes: [
      { path: '/', component: LoginView }
      // { path: '/new/:page(\\d+)?', component: createListView('new') },
      // { path: '/show/:page(\\d+)?', component: createListView('show') },
      // { path: '/ask/:page(\\d+)?', component: createListView('ask') },
      // { path: '/job/:page(\\d+)?', component: createListView('job') },
      // { path: '/item/:id(\\d+)', component: ItemView },
      // { path: '/user/:id', component: UserView },
      // { path: '/', redirect: '/login' }
    ]
  });
}
