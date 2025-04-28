<template>
  <div>
    <h1 m-1 c-1 animate-fade-in-left font-allura>196.32</h1>
    <n-card title="小卡片" size="small" hoverable :class="{active: value > 0}">
      <n-input-number
        v-model:value="value"
        :min="0"
        :max="99"
        :precision="0"
        :bordered="false"
        size="tiny"
        button-placement="both"
        @update:value="numberInput"
        w-20
      />
    </n-card>
  </div>
</template>

<script setup>
import { ref } from "vue";

let value = ref(0);

const numberInput = (v) => {
  console.log(v);
};
</script>

<style scoped>
.active {
    border: 1px solid black;
}
</style>
