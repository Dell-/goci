import {USER_LOGIN, USER_LOGOUT, USER_REFRESH} from '../mutation-types';
import Vue from 'vue';

const state = {
  isAuthenticated: false
};

const mutations = {
  [USER_LOGIN](state) {
    state.isAuthenticated = true;
  },
  [USER_LOGOUT](state) {
    state.isAuthenticated = false;
  },
  [USER_REFRESH](state, data) {

  }
};

const actions = {
  login({commit}, user) {
    return new Promise((resolve) => {
      Vue.auth.login(user)
        .then(() => {
          commit(USER_LOGIN);
          resolve();
        });
    });
  },
  logout({commit}) {
    return new Promise((resolve) => {
      Vue.auth.logout()
        .then(() => {
          commit(USER_LOGOUT);
          resolve();
        });
    });
  }
};

export default {
  state,
  mutations,
  actions
};
