import Vue from 'vue'
import Router from 'vue-router'
import index from '@/components/index'
import notification from '@/components/services/notification.vue'
import alert from "vue-simple-alert";

Vue.use(Router)
Vue.use(alert);

export default new Router({
  mode: 'history',
  routes: [
    { path: '/', name: 'Home', component: index },
    { path: '/services/notification', name: 'Notification Service', component: notification }
  ]
})
