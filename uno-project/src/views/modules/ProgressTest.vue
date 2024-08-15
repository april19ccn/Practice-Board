<template>
    <div style="width: 500px" mt-10>
        <n-list hoverable clickable>
            <n-list-item v-for="(item, index) in data" :key="index">
                <div class="flex flex-items-center">
                    <div class="w-5 flex-shrink-0 color-orange font-bold">{{ index + 1 }}</div>

                    <n-tooltip placement="bottom" :show-arrow="false" trigger="hover">
                        <template #trigger>
                            <div class="w-30 flex-shrink-0 text-truncate">{{ item.label }}</div>
                        </template>
                        <span> {{ item.label }} </span>
                    </n-tooltip>

                    <n-progress :id="getProId(item.percentage)" class="ml-2" type="line" :percentage="item.percentage" :show-indicator="false" />

                    <div class="flex flex-items-baseline flex-justify-end ml-5 w-30">
                        <p class="per-text">{{ parseFloat(item.percentage.toFixed(1)) }}</p>
                        <p class="per-sign" ml-1>%</p>
                    </div>
                </div>
            </n-list-item>
        </n-list>
    </div>
</template>

<script setup>
import { ref } from "vue";

const data = ref([
    {
        label: "长盈终有别,长盈终有别,长盈终有别",
        percentage: 18,
    },
    {
        label: "秋风渐入深",
        percentage: 75.99,
    },
    {
        label: "长盈终有别",
        percentage: 45.2,
    },
]);

const getProId = (percentage) => {
    if (percentage >= 0 && percentage < 30) {
        return "progress-color";
    }

    if (percentage >= 30 && percentage < 60) {
        return "progress-color-1";
    }

    if (percentage >= 60 && percentage < 100) {
        return "progress-color-2";
    }
};
</script>

<style scoped>
#progress-color:deep(.n-progress-graph-line-fill) {
    background: linear-gradient(161deg, #f6d242, #ff52e5);
}

#progress-color-1:deep(.n-progress-graph-line-fill) {
    background: linear-gradient(71deg, #f54ea2, #ff7676);
}

#progress-color-2:deep(.n-progress-graph-line-fill) {
    background: linear-gradient(71deg, #fccf31, #f55555);
}

.per-text {
    font-size: 22px;
    font-weight: 600;
    color: #535353;
}

.per-sign {
    font-size: 13px;
    font-weight: bold;
    color: #9c9c9c;
}
</style>
