<script setup lang="ts">
import {ref} from 'vue'
// import TheWelcome from '../components/TheWelcome.vue'
const handleScroll = () => {
  // 在这里处理滚动事件
  // console.log('滚动事件被触发', event);
}

// 处理滚动事件
let startTouchTime = ref(0)
let isScrolling = ref(false)
const handleTouchStart = () => {
    startTouchTime.value = new Date().getTime();
    isScrolling.value = false;
}
const handleTouchMove = () => {
    isScrolling.value = true;
}
const handleTouchEnd = (index: any) => {
    if (!isScrolling.value) {
        const duration = new Date().getTime() - startTouchTime.value;
        console.log(duration)
        // 设置一个阈值，比如100ms以下认为是点击
        if (duration < 150) {
            // 触控点击事件处理
            console.log("touch-click" + index)
        }
    }
    isScrolling.value = false
}
</script>

<template>
  <main>
    <!-- <TheWelcome /> -->
    {{ isScrolling }}
    <br/>
    {{ startTouchTime }}
    <br/>
    {{ new Date().getTime() - startTouchTime }}
    <div class="scroll-container" @scroll="handleScroll">
      <!-- 滚动内容 -->
      <div class="scroll-content" 
        @touchstart="handleTouchStart"
        @touchmove="handleTouchMove" 
        @touchend="handleTouchEnd(1)">
        <!-- 滚动的内容 -->
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
        <p>1</p>
      </div>
    </div>
  </main>
</template>

<style>
.scroll-container {
  overflow-y: scroll; /* 设置容器为垂直滚动 */
  height: 200px; /* 设置容器高度 */
  width: 400px; /* 设置容器宽度 */
  background-color: aqua;
}
.scroll-content {
  height: 600px; /* 设置内容高度，超过容器高度以便滚动 */
  width: 100%;
  background-color: darkorange;
}

.scroll-container::-webkit-scrollbar {
    width: 8px;
}

.scroll-container::-webkit-scrollbar-track {
    box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
    border-radius: 10px;
    background: #ededed;
}

.scroll-container::-webkit-scrollbar-thumb {
    background: #838383;
    border-radius: 4px;
}

.scroll-container::-webkit-scrollbar-thumb:hover {
    background: #6e6e6e;
}

.scroll-container::-webkit-scrollbar-thumb:active {
    background: #575757;
}
</style>
