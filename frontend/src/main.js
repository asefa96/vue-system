// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import { store } from "./store/store"
import {
  faCogs,
  faEthernet,
  faHome, faMemory, faMicrochip, faNetworkWired, faServer,
 
} from '@fortawesome/free-solid-svg-icons'
import { library as fontAwesomeLib } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

fontAwesomeLib.add([
  faHome,
  faMicrochip,
  faNetworkWired,
  faMemory,
  faCogs,
  faServer
  
  


])

Vue.component('fa', FontAwesomeIcon)

Vue.config.productionTip = false
Vue.prototype.$PORT=8888
/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,
  components: { App },
  template: '<App/>'
})
