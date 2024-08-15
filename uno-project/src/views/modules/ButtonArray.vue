<template>
  <div>
    <n-space>
      <n-button
        v-for="item of arr"
        :key="item.key"
        type="primary"
        :ghost="setActive != item.key"
        @click="handle(item)"
      >
        {{ item.label }}
      </n-button>
    </n-space>
  </div>
</template>

<script setup>
import { ref, reactive } from "vue";

let arr = reactive([
  {
    key: 1,
    label: "all",
  },
  {
    key: 2,
    label: "test",
  },
]);

const setActive = ref(1);
const handle = (item) => {
  if (setActive.value == item.key) return;
  console.log(item);
  setActive.value = item.key;
};
</script>

<style scoped></style>
