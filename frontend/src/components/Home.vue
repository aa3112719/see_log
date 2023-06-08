<template>
  <a-layout class="layout">
    <a-layout-header>
      <a-button class="btn" @click="fileInfo">选择文件</a-button>
    </a-layout-header>
    <a-layout style="padding: 0 24px">
      <a-layout-content>
        <a-typography-text :type="formatTypeByTxt(item)" v-for="item in result">
          {{ item }}
        </a-typography-text>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const result = ref([])
const name = ref("")

const greet = () => {
  window.go.main.App.Greet(name.value).then(response => {
    result.value = response
    console.log(response)
  })
}

const fileInfo = () => {
  window.go.main.App.FileInfo().then((res: any) => {
    result.value = res
  }).catch((err: any) => {
    console.log(err)
  })
}

const formatTypeByTxt = (txt: string) => {
  if (txt.indexOf('error') !== -1) return 'danger'
  if (txt.indexOf('info') !== -1) return 'success'
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
</style>
