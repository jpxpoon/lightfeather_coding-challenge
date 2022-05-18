<template>
  <div id="app">
    <navigation></navigation>
    <router-view/>
  </div>
</template>

<script>
import navigation from './components/nav.vue'
import notification from './components/services/notification'
export default {
  name: 'App',
   components: {
    navigation,
    notification
  }
}
</script>