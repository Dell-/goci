<template>
    <div id="app" class="app">
        <auth-layout v-if="!isAuthenticated"></auth-layout>
        <layout v-else></layout>
    </div>
</template>

<script>
  import Layout from 'components/layout/Layout';
  import AuthLayout from './components/layout/AuthLayout';
  import VuesticPreLoader from './components/vuestic-components/vuestic-preloader/VuesticPreLoader.vue';
  import {mapGetters} from 'vuex';

  export default {
    name: 'app',
    components: {
      VuesticPreLoader,
      AuthLayout,
      Layout
    },
    computed: mapGetters({
      isAuthenticated: 'isAuthenticated'
    })
  };
</script>

<style lang="scss">
    @import "sass/main";

    body {
        height: 100%;
        .app {
            height: 100%;
        }
    }
</style>
