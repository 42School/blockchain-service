// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import VueSession from 'vue-session'
import VueParticles from 'vue-particles'

Vue.use(VueSession)
Vue.use(VueParticles)

require('./scss/styles.sass');

Vue.config.productionTip = false
Vue.prototype.$http = axios

axios.defaults.baseURL = process.env.BACK_URL
axios.defaults.headers.common['Accept'] = 'application/json'
axios.defaults.headers.common['Content-Type'] = 'application/json; charset=utf-8'
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*'
axios.defaults.headers.withCredentials = true
axios.defaults.headers.credentials = 'same-origin'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
