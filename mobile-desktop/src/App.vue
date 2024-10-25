<script setup>
import DownLoad from './components/DownLoad.vue';

const resizeHandler = () => {
  const rect = document.body.getBoundingClientRect()
  const isMobile = rect.width < 980
  appStore.setDevice(isMobile ? 'mobile' : 'desktop')
  appStore.setCollapse(isMobile)
}

</script>

<template>
  <main class="main-content">
    <DownLoad />
  </main>
</template>

<style scoped>
.main-content {
    height: 100vh;
    display: flex;
    flex-direction: column;
}
</style>
