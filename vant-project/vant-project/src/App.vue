<script setup lang="ts">
import { ref } from "vue";
import { SolarDay, LunarYear } from "tyme4ts";

//
const solar = SolarDay.fromYmd(2024, 9, 5);
console.log(solar.toString());
console.log(solar.getLunarDay().toString());

//
let lunarYear = LunarYear.fromYear(2024);
let months = lunarYear.getMonths();
console.log(months);

const MONTHIS_LIST = [{text: "正月", value: "Monday"}, {text: "二月", value: "Tuesday"}, {text: "三月", value: "Wednesday"}, {text: "四月", value: "Thursday"}, {text: "五月", value: "Friday"}, {text: "六月", value: "Saturday"}, {text: "七月", value: "Sunday"}, {text: "八月", value: "Monday"}, {text: "九月", value: "Tuesday"}, {text: "十月", value: "Wednesday"}, {text: "十一月", value: "Thursday"}, {text: "十二月", value: "Friday"}];
const DAY_LIST = [
    { text: "一", value: "1" },
    { text: "二", value: "2" },
    { text: "三", value: "3" },
    { text: "四", value: "4" },
    { text: "五", value: "5" },
    { text: "六", value: "6" },
    { text: "七", value: "7" },
    { text: "八", value: "8" },
    { text: "九", value: "9" },
]

const columns = ref([]);
const loading = ref(true);

setTimeout(() => {
    columns.value = [
        // 第一列
        [
            { text: "周一", value: "Monday" },
            { text: "周二", value: "Tuesday" },
            { text: "周三", value: "Wednesday" },
            { text: "周四", value: "Thursday" },
            { text: "周五", value: "Friday" },
        ],
    ];

    columns.value.push(MONTHIS_LIST)
    columns.value.push(DAY_LIST)

    loading.value = false;
}, 1000);

const currentData = ref()
</script>

<template>
    {{ currentData }}
    <van-picker v-model="currentData" :columns="columns" :loading="loading" />
</template>

<style scoped></style>
