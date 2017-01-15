import Vue from 'vue'
import VueResource from 'vue-resource'
import { sync } from 'vuex-router-sync'

import store from './store'
import router from './router'
import App from './app.vue'
import * as filters from './filters'

Vue.use(VueResource)

sync(
  store,
  router
)

Object.keys(filters).forEach((key) => {
  Vue.filter(key, filters[key])
})

new Vue({
  http: {
    root: window.INITIAL_STATE['api']
  },

  el: '#app',
  store,
  router,
  render: h => h(App)
})
