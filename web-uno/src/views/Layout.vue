<template>
    <h1 text-very-cool>test</h1>
    <button @click="show = !show">Toggle Slide + Fade</button>
    <n-layout has-sider h-screen>
        <n-layout>
            <n-layout-content h-screen>
                <ImgCarousel></ImgCarousel>
            </n-layout-content>
        </n-layout>
        <n-layout-sider content-style="padding: 24px;" :width="450" bgc-w>
            <div flex items-center h-full>
                <Transition name="slide-fade">
                    <div flex-1 v-if="show">
                        <div>
                            <n-gradient-text type="error" gradient="linear-gradient(135deg,#7117ea,#ea6060)">
                                <span font-size-8>测试UNO</span>
                            </n-gradient-text>

                            <n-gradient-text type="error" gradient="linear-gradient(135deg,#7117ea,#ea6060)" ml-2>
                                <span font-allura font-size-6>Test UI Layout</span>
                            </n-gradient-text>
                        </div>

                        <n-icon size="40" color="#0e7a0d">
                            <game-controller />
                        </n-icon>
                        <div>
                            <n-form ref="formRef" :label-width="80" :model="formValue" :rules="rules" size="large">
                                <n-form-item path="user.name">
                                    <template #label>
                                        <n-icon size="40" color="#0e7a0d">
                                            <game-controller />
                                        </n-icon>
                                    </template>
                                    <n-input v-model:value="formValue.user.name" placeholder="输入姓名" />
                                </n-form-item>
                                <n-form-item label="年龄" path="user.age">
                                    <n-input v-model:value="formValue.user.age" placeholder="输入年龄" />
                                </n-form-item>
                                <n-form-item label="电话号码" path="phone">
                                    <n-input v-model:value="formValue.phone" placeholder="电话号码" />
                                </n-form-item>
                                <n-form-item>
                                    <n-button attr-type="button" @click="handleValidateClick">
                                        验证
                                    </n-button>
                                </n-form-item>
                            </n-form>
                        </div>
                    </div>
                </Transition>
            </div>
        </n-layout-sider>
    </n-layout>
</template>

<script setup lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import { type FormInst } from 'naive-ui'

import ImgCarousel from '@/components/ImgCarousel.vue'

const show = ref(false)

const formRef = ref<FormInst | null>(null)

const formValue = ref({
    user: {
        name: '',
        age: ''
    },
    phone: ''
})

const rules = ref({
    user: {
        name: {
            required: true,
            message: '请输入姓名',
            trigger: 'blur'
        },
        age: {
            required: true,
            message: '请输入年龄',
            trigger: ['input', 'blur']
        }
    },
    phone: {
        required: true,
        message: '请输入电话号码',
        trigger: ['input']
    }
})


const handleValidateClick = (e: MouseEvent) => {
    e.preventDefault()
    formRef.value?.validate((errors) => {
        if (!errors) {
            (window as any).$message.success('Valid')
        } else {
            // console.log(errors)
            (window as any).$message.error('Invalid')
        }
    })
}

onMounted(() => {
    show.value = true
})
</script>

<style scoped>
.n-layout-header,
.n-layout-footer {
    background: rgba(128, 128, 128, 0.2);
    padding: 24px;
}

.n-layout-sider {
    background: rgba(128, 128, 128, 0.3);
}

.n-layout-content {
    background: rgba(128, 128, 128, 0.4);
}

.slide-fade-enter-active {
    transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
    transform: translateX(20px);
    opacity: 0;
}
</style>