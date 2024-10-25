<template>
    <div class="header">
        <h1>Get Started</h1>
    </div>

    <div class="download-content">
        <Transition name="slide-fade">
            <WelcomeItem v-if="show1">
                <template #icon>
                    <DocumentationIcon />
                </template>

                <template #heading>Documentation</template>

                Vue’s <span>official documentation</span> provides you with all information you need to get started.
            </WelcomeItem>
        </Transition>

        <Transition name="slide-fade">
            <WelcomeItem v-if="show2">
                <template #icon>
                    <ToolingIcon />
                </template>

                <template #heading>ToolingIcon</template>

                Vue’s <span>official documentation</span> provides you with all information you need to get started.
            </WelcomeItem>
        </Transition>

        <Transition name="slide-fade">
            <WelcomeItem v-if="show3">
                <template #icon>
                    <EcosystemIcon />
                </template>

                <template #heading>EcosystemIcon</template>

                Vue’s <span>official documentation</span> provides you with all information you need to get started.
            </WelcomeItem>
        </Transition>

        <Transition name="slide-fade">
            <WelcomeItem v-if="show4">
                <template #icon>
                    <SupportIcon />
                </template>

                <template #heading>SupportIcon</template>

                Vue’s <span>official documentation</span> provides you with all information you need to get started.
            </WelcomeItem>
        </Transition>
    </div>

    <Transition name="button-fade">
        <button v-if="show1" class="download-btn">Get Started</button>
    </Transition>
</template>

<script setup>
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from '../icons/IconDocumentation.vue'
import ToolingIcon from '../icons/IconTooling.vue'
import EcosystemIcon from '../icons/IconEcosystem.vue'
import SupportIcon from '../icons/IconSupport.vue'

import { ref, onBeforeMount } from 'vue'

const show1 = ref(false)
const show2 = ref(false)
const show3 = ref(false)
const show4 = ref(false)

onBeforeMount(() => {
    setTimeout(() => {
        show1.value = true
    }, 300)
    setTimeout(() => {
        show2.value = true
    }, 500)
    setTimeout(() => {
        show3.value = true
    }, 700)
    setTimeout(() => {
        show4.value = true
    }, 900)
})
</script>

<style scoped>
span {
    color: #f8a2c8;
    font-weight: bold;
}

.header {
    padding: 0 2rem;
    padding-top: 60px;
    padding-bottom: 10px;
}

.download-content {
    padding: 0 2rem;
    padding-bottom: 8rem;
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.download-btn {
    width: 80vw;
    height: 50px;
    position: fixed;
    bottom: 40px;
    left: 0;
    right: 0;
    margin: auto;
    border: 1px solid #ddc1ce;
    background-color: #fff;
    border-radius: 5px;
}

/*
  进入和离开动画可以使用不同
  持续时间和速度曲线。
*/
.slide-fade-enter-active {
    transition: all 0.3s cubic-bezier(1, 0.5, 0.6, 1);
}

.slide-fade-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
    transform: translateX(25px);
    opacity: 0;
}

/*
  进入和离开动画可以使用不同
  持续时间和速度曲线。
*/
.button-fade-enter-active {
    transition: all 0.3s cubic-bezier(1, 0.5, 0.6, 1);
}

.button-fade-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.button-fade-enter-from,
.button-fade-leave-to {
    transform: translateY(20px);
    opacity: 0;
}
</style>