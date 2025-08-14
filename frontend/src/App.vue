<template>
  <div id="app">
    <Navigation />
    <main class="main-content">
      <router-view />
    </main>
  </div>
</template>

<script>
import Navigation from '@/components/Navigation.vue'
import { useAuthStore } from '@/stores/auth'

export default {
  name: 'App',
  components: {
    Navigation
  },
  mounted() {
    const authStore = useAuthStore()
    authStore.initAuth()
  }
}
</script>

<style>
#app {
  min-height: 100vh;
  background: linear-gradient(135deg, var(--cherry-bg) 0%, #fff 100%);
}

.main-content {
  padding-top: 80px; /* Space for fixed navigation */
  min-height: calc(100vh - 80px);
}
</style>