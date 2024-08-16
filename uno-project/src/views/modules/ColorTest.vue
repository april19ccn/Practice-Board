<template>
    <div>
        <div class="theme-element">主题色变换示例</div>


        <div class="box">
             <div class="text">渐变动画</div>
        </div>
    </div>
</template>

<script setup lang="ts"></script>

<style scoped>
/* 定义关键帧动画 */
@keyframes theme-color-change {
    0% {
        background-color: blue; /* 初始颜色 */
    }
    50% {
        background-color: red; /* 结束颜色 */
    }
    100% {
        background-color: blue; /* 初始颜色 */
    }
}

/* 应用动画到需要变换颜色的元素 */
.theme-element {
    animation: theme-color-change 5s infinite; /* 持续时间5秒，并且无限循环 */
}

.box{
    /* 初始化 取消内外边距 */
    /* margin: 0;
    padding: 0; */
    /* 100%窗口高度 */
    /* height: 100vh; */

    width: 200px;
    height: 200px;
    /* 弹性布局 水平垂直居中 */
    display: flex;
    justify-content: center;
    align-items: center;
    /* 渐变背景 */
    background: linear-gradient(125deg,#2c3e50,#27ae60,#2980b9,#e74c3c,#8e44ad);
    /* 指定背景图像的大小 */
    background-size: 500%;
    /* 执行动画：动画名 时长 线性的 无限次播放 */
    animation: bgAnimation 15s linear infinite;

    filter: blur(100px);
    border-radius: 50%;
}
.text{
    color: #fff;
    font-size: 30px;
    /* 字间距 */
    letter-spacing: 15px;
}

/* 定义动画 */
@keyframes bgAnimation {
    0%{
        background-position: 0% 50%;
    }
    50%{
        background-position: 100% 50%;
    }
    100%{
        background-position: 0% 50%;
    }
}
</style>
