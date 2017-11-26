const AUTH_KEY = 'auth-token';

const API_LOGIN = 'auth/login';
const API_LOGOUT = 'auth/logout';
const API_REGISTRATION = 'users';
const API_CURRENT_USER = 'users/current';

export const Auth = {
  install(Vue) {
    const service = new AuthService(Vue.http, API_LOGIN, API_LOGOUT, API_REGISTRATION, API_CURRENT_USER);
    Vue.auth = service;
    Object.defineProperties(Vue.prototype, {
      $auth: {
        get: () => {
          return service;
        }
      }
    });
  }
};

class AuthService {
  constructor(http, login, logout, registration, current) {
    this.http = http;
    this.loginUrl = login;
    this.logoutUrl = logout;
    this.registrationUrl = registration;
    this.currentUrl = current;
  }

  login(data) {
    return this.http.post(this.loginUrl, data)
      .then((r) => setToken(JSON.stringify(r.body)));
  }

  logout() {
    return this.http.post(this.logoutUrl).then(removeToken);
  }

  registration(data) {
    return this.http.post(this.registrationUrl, data);
  }

  current(data) {
    return this.http.get(this.currentUrl);
  }

  token() {
    return getToken() || false;
  }
}

const setToken = (token) => window.localStorage.setItem(AUTH_KEY, token);
const removeToken = () => window.localStorage.removeItem(AUTH_KEY);
const getToken = () => JSON.parse(window.localStorage.getItem(AUTH_KEY));
