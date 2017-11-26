<template>
    <div class="login">
        <h2>Welcome!</h2>
        <form v-on:submit.prevent="login" method="post" action="/auth/login" name="login">
            <div class="form-group">
                <div class="input-group">
                    <input v-model="email" type="text" id="email" required="required"/>
                    <label class="control-label" for="email">Email</label><i class="bar"></i>
                </div>
            </div>
            <div class="form-group">
                <div class="input-group">
                    <input v-model="password" type="password" id="password" required="required"/>
                    <label class="control-label" for="password">Password</label><i class="bar"></i>
                </div>
            </div>
            <div class="d-flex flex-column flex-lg-row align-items-center justify-content-between down-container">
                <button class="btn btn-primary" type="submit">
                    Log In
                </button>
                <router-link class='link' :to="{name: 'Signup'}">Create account</router-link>
            </div>
        </form>
    </div>
</template>

<script>
  import {ROUTE_DASHBOARD} from './../../../router/path';

  const user = {
    email: 'test@test.loc',
    password: '123123q'
  };

  export default {
    name: 'login',
    data() {
      return user;
    },
    methods: {
      login() {
        this.$store.dispatch('login', user)
          .then(() => this.$router.replace(ROUTE_DASHBOARD));
      },
      register() {

      }
    }
  };
</script>

<style lang="scss">
    @import '../../../sass/variables';
    @import '../../../../node_modules/bootstrap/scss/mixins/breakpoints';
    @import '../../../../node_modules/bootstrap/scss/variables';

    .login {
        @include media-breakpoint-down(md) {
            width: 100%;
            padding-right: 2rem;
            padding-left: 2rem;
            .down-container {
                .link {
                    margin-top: 2rem;
                }
            }
        }

        h2 {
            text-align: center;
        }
        width: 21.375rem;

        .down-container {
            margin-top: 3.125rem;
        }
    }
</style>
