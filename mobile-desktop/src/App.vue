<script setup>
import DownLoad from './components/DownLoad.vue';
import { ref,onBeforeMount, onBeforeUnmount } from 'vue';

const isMobile = ref(false)

const resizeHandler = () => {
  const rect = document.body.getBoundingClientRect()
  const isMobile = rect.width < 980
}

// vueuse/core
const debounceFn = useDebounceFn(resizeHandler, 100)

onBeforeMount(() => {
  resizeHandler()
  window.addEventListener('resize', debounceFn)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', debounceFn)
})
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
