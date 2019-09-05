import Vue from 'vue'
import Router from 'vue-router'
import index from '@/pages/index'
import search from '@/pages/search'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: index
    },
    {
      path: '/search/:login?',
      name: 'search',
      component: search
    }
  ],
});
