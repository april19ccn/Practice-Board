<template>
  <div class="about">
    <h1 m-1 c-1 animate-fade-in-left font-allura>This is an about page</h1>

    <n-config-provider :theme="darkTheme">
      <n-card title="卡片" hoverable>
        <n-gradient-text type="error" gradient="linear-gradient(135deg,#7117ea,#ea6060)">
          <span font-size-10>测试</span>
        </n-gradient-text>

        <n-gradient-text type="error" gradient="linear-gradient(135deg,#7117ea,#ea6060)" ml-2>
          <span font-allura font-size-12>This is an about page</span>
        </n-gradient-text>
      </n-card>
    </n-config-provider>

    <ImgCarousel></ImgCarousel>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { darkTheme } from 'naive-ui'
import ImgCarousel from '@/components/ImgCarousel.vue'

export default defineComponent({
  setup() {
    return {
      darkTheme
    }
  }
})
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
