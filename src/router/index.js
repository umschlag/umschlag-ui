import Vue from 'vue';
import VueRouter from 'vue-router';

import DashboardView from '../views/dashboard.vue';
import ProfileView from '../views/profile.vue';

Vue.use(VueRouter);

export default new VueRouter({
  mode: 'history',
  base: __dirname,
  scrollBehavior: () => ({ y: 0 }),
  routes: [
     { path: '/profile', component: ProfileView },
     { path: '/', component: DashboardView },
  ],
});
