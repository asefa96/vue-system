import Vue from 'vue'
import Router from 'vue-router'
import SystemInfo from '@/views/SystemInfo'

Vue.use(Router)

export default new Router({
    routes: [{
        path: '/',
        name: 'SystemInfo',
        component: SystemInfo
    }]
})