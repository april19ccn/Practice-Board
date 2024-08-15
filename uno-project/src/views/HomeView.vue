<template>
  <div>
    <n-space>
      <ButtonArray></ButtonArray>
      <CardArray></CardArray>
      <LayoutTest></LayoutTest>
    </n-space>
    <n-space>
        <ProgressTest></ProgressTest>
    </n-space>
  </div>
</template>

<script setup>
import ButtonArray from './modules/ButtonArray.vue'
import CardArray from './modules/CardArray.vue'
import LayoutTest from './modules/LayoutTest.vue';
import ProgressTest from './modules/ProgressTest.vue';
</script>

<style scoped></style>
