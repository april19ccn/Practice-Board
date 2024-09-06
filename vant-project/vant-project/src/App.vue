<script setup lang="ts">
import { onMounted, ref } from "vue";
import { SolarDay, LunarYear } from "tyme4ts";

//
var now = new Date();
// const today = SolarDay.fromYmd(now.getFullYear(), now.getMonth() + 1, now.getDate());
const today = SolarDay.fromYmd(now.getFullYear(), 1, now.getDate());
console.log(today.toString());
console.log(today.getLunarDay().toString());

//
let lunarYear = LunarYear.fromYear(2024);
let months = lunarYear.getMonths();
console.log(months);

const MONTHIS_LIST = [
    { text: "正月", value: "正" },
    { text: "二月", value: "二" },
    { text: "三月", value: "三" },
    { text: "四月", value: "四" },
    { text: "五月", value: "五" },
    { text: "六月", value: "六" },
    { text: "七月", value: "七" },
    { text: "八月", value: "八" },
    { text: "九月", value: "九" },
    { text: "十月", value: "十" },
    { text: "十一月", value: "十一" },
    { text: "腊月", value: "腊" },
];
const DAY_LIST = [
    { text: "一", value: 1 },
    { text: "二", value: 2 },
    { text: "三", value: 3 },
    { text: "四", value: 4 },
    { text: "五", value: 5 },
    { text: "六", value: 6 },
    { text: "七", value: 7 },
    { text: "八", value: 8 },
    { text: "九", value: 9 },
];

const columns = ref();
const loading = ref(true);

const init = () => {
    // YEAR
    columns.value = [[]];
    for (let i = 0; i < 100; i++) {
        columns.value[0].push({ text: 2024 - i, value: 2024 - i });
    }

    // MONTH
    columns.value.push(MONTHIS_LIST);

    // DAY
};

setTimeout(() => {

    columns.value.push(DAY_LIST);

    loading.value = false;
}, 1000);

// 
const currentData = ref();

onMounted(() => {
    init()
})
</script>

<template>
    {{ currentData }}
    <van-picker v-model="currentData" :columns="columns" :loading="loading" />
</template>

<style scoped></style>
