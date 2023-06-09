<template>
    <a-layout class="layout">
        <a-layout-header style="padding: 24px">
            <a-row :gutter="2">
                <a-col :xl="6" :xs="6">
                    <a-button style="margin-bottom: 5px;" @click="fileInfo">选择文件</a-button>
                </a-col>
                <a-col :xl="16" :xs="14">
                    <a-typography-paragraph
                        style="
                            width: 100%;
                            overflow: hidden;
                            white-space: nowrap;
                            text-overflow: ellipsis;
                        "
                        :ellipsis="{ rows: 1 }"
                    >
                        {{ selection }}
                    </a-typography-paragraph>
                </a-col>
                <a-col :span="2"><a-spin v-if="loading" /></a-col>
            </a-row>
            <a-input v-model="highLightStr" placeholder="请输入需要高亮的字符串" @press-enter="handleScrollTo"/>
        </a-layout-header>
        <a-layout style="padding: 0 24px">
            <a-layout-content>
                <a-scrollbar
                    ref="scrollRef"
                    style="height: 600px; width: 100%; overflow: auto"
                    :outer-style="{ 'max-height': '1200px' }"
                >
                    <a-space direction="vertical" :size="10">
                        <a-typography-text
                            :type="formatTypeByTxt(item)"
                            v-for="(item, index) in result"
                        >
                            <a-space>
                                <div>【{{ index }}】</div>
                                <div style="word-break: break-all" v-html="formatHighLightStr(item)"></div>
                            </a-space>
                        </a-typography-text>
                    </a-space>
                </a-scrollbar>
            </a-layout-content>
        </a-layout>
    </a-layout>
</template>

<script setup lang="ts">
import { onUnmounted, ref } from 'vue'
import useLoading from '../hooks/useLoading'

const result = ref<Array<string>>([])
const name = ref('')
const selection = ref('')

const greet = () => {
    window.go.main.App.Greet(name.value).then((response) => {
        result.value = response
        console.log(response)
    })
}
let fileInfoTimer = ref()

const { loading, setLoading } = useLoading()

const fileInfo = () => {
    window.go.main.App.FileInfo()
        .then((res: any) => {
            selection.value = res
            if (!!res) {
                setLoading(true)
                fileInfoTimer.value = setInterval(() => {
                    getFileStr()
                }, 2000)
            }
        })
        .catch((err: any) => {
            console.log(err)
        })
}

const getFileStr = () => {
    window.go.main.App.GetFileStr(logNum.value)
        .then((res: Array<string>) => {
            result.value = res
        })
        .finally(() => {
            setLoading(false)
        })
}

const logNum = ref(100)

onUnmounted(() => {
    clearInterval(fileInfoTimer.value)
})

const formatTypeByTxt = (txt: string) => {
    if (txt.indexOf('[error]') !== -1) return 'danger'
    if (txt.indexOf('[info]') !== -1) return 'success'
}

const highLightStr = ref('')

const formatHighLightStr = (str: string) => {
    const toolMan = str.replace(highLightStr.value, '~(*)~')
    return toolMan.replace(
        '~(*)~',
        `<span class="color-high-light">${highLightStr.value}</span>`
    ).replace('<span class="color-high-light"></span>', '')
}

const scrollRef = ref(null)
const scrollToNum = ref(0)
const handleScrollTo = () => {
  const el = document.getElementsByClassName('color-high-light')[scrollToNum.value]

  if(!el) {
    const zero = document.getElementsByClassName('color-high-light')[0]
    if(zero) {
      scrollToNum.value = 0
    }
  }

  const ele = document.getElementsByClassName('color-high-light')[scrollToNum.value]
  const distanceToTop = (ele as any).offsetTop;
  (scrollRef.value as any).scrollTop(distanceToTop)

  if(document.getElementsByClassName('color-high-light')[scrollToNum.value + 1]) {
    scrollToNum.value = scrollToNum.value  + 1
  } else {
    scrollToNum.value = 0
  }
}

</script>

<style>
.layout {
    width: 100%;
    height: 100%;
}

.logos {
    display: flex;
    justify-content: space-around;
    align-items: center;
    height: 50%;
    width: 100%;
}

.logos img.wails {
    max-width: 256px;
    max-height: 256px;
}

.logos img.vue {
    max-width: 256px;
    max-height: 185px;
}

.links {
    font-size: 20px;
}

.result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
}

.input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
}

.input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
}

.input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
}

.layout-content {
    min-width: 1024px;
    min-height: 100vh;
    overflow-y: hidden;
    background-color: var(--color-fill-2);
    transition: all 0.2s cubic-bezier(0.34, 0.69, 0.1, 1);
}
.color-high-light {
    color: #ab0505;
    background-color: #ffb500;
}
</style>
