<template>
  <div :class="classObject" v-resize>
    <navbar></navbar>
    <div class="content-wrap" id="content-wrap">
      <main id="content" class="content" role="main">
        <vuestic-pre-loader v-show="isLoading" ref="preLoader" class="pre-loader"></vuestic-pre-loader>
        <router-view v-show="!isLoading"></router-view>
      </main>
    </div>
  </div>
</template>

<script>
  import {mapGetters} from 'vuex';

  import Navbar from './navbar/Navbar';
  import Resize from 'directives/ResizeHandler';
  import VuesticPreLoader from '../vuestic-components/vuestic-preloader/VuesticPreLoader';

  export default {
    name: 'layout',

    components: {
      Navbar,
      VuesticPreLoader
    },
    directives: {
      resize: Resize
    },
    computed: {
      ...mapGetters([
        'toggleWithoutAnimation',
        'isLoading'
      ]),
      classObject: function() {
        return {
          'sidebar-hidden': !this.toggleWithoutAnimation,
          'sidebar-hidden sidebar-hidden_without-animation': this.toggleWithoutAnimation
        };
      }
    }
  };
</script>

<style lang="scss">
  @import "../../sass/_variables.scss";
  @import "../../../node_modules/bootstrap/scss/mixins/breakpoints";
  @import "../../../node_modules/bootstrap/scss/variables";

  .content-wrap {
    margin-left: $content-wrap-ml;
    padding: $content-wrap-pt $content-wrap-pr $content-wrap-pb 0;
    transition: margin-left 0.3s ease;

    .pre-loader {
      position: absolute;
      left: $vuestic-preloader-left;
      top: $vuestic-preloader-top;
    }


    @include media-breakpoint-down(md) {
      padding: $content-mobile-wrap;
      margin-left: 0;
    }
  }
</style>
