import Vue from 'vue';
import VueResource from 'vue-resource';
import { sync } from 'vuex-router-sync';

import store from './store';
import router from './router';
import App from './app.vue';
import * as filters from './filters';

Vue.use(VueResource);

sync(
  store,
  router,
);

Object.keys(filters).forEach((key) => {
  Vue.filter(key, filters[key]);
});

new Vue({
  // http: {
  //   root: '/root',
  //   headers: {
  //     Authorization: 'Basic YXBpOnBhc3N3b3Jk'
  //   }
  // },

  el: '#app',
  store,
  router,
  render: h => h(App),
});
