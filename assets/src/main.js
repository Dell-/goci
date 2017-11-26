// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import BootstrapVue from 'bootstrap-vue';
import VeeValidate from 'vee-validate';
import App from './App';
import store from './store';
import router from './router';
import {sync} from 'vuex-router-sync';
import VueResource from 'vue-resource';
import {Auth} from './plugins/auth';

import {API_URL} from './config';

import {ROUTE_AUTH_LOGIN} from './router/path';

Vue.use(BootstrapVue);

// NOTE: workaround for VeeValidate + vuetable-2
Vue.use(VeeValidate, {fieldsBagName: 'formFields'});

/**
 * Documentation for Vue Resource
 * https://github.com/pagekit/vue-resource
 */
Vue.use(VueResource);
Vue.http.options.root = API_URL;

Vue.use(Auth);

sync(store, router);

let mediaHandler = () => {
  if (window.matchMedia(store.getters.config.windowMatchSizeLg).matches) {
    store.dispatch('toggleSidebar', true);
  } else {
    store.dispatch('toggleSidebar', false);
  }
};

router.beforeEach((to, from, next) => {
  let token = Vue.auth.token();
  if (token) {
    Vue.http.headers.common['Authorization'] = 'Bearer ' + token.token;
  }

  next();
});

router.beforeEach((to, from, next) => {
  const authRequired = to.matched.some(record => record.meta.auth);
  const isAuthenticated = store.getters.isAuthenticated;

  if (authRequired && !isAuthenticated) {
    next(ROUTE_AUTH_LOGIN);
  } else if (isAuthenticated && to.path === ROUTE_AUTH_LOGIN) {
    next(false);
  } else {
    next();
  }
});

router.beforeEach((to, from, next) => {
  store.commit('setLoading', true);
  next();
});

router.afterEach((to, from) => {
  mediaHandler();
  store.commit('setLoading', false);
});

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: {App}
});
