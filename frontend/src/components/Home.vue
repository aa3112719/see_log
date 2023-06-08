<template>
  <div style="display: flex; justify-content: center; align-items: center; height: 100vh; color: white">
    <a-list>
      <a-list-item v-for="item in result">
        {{ item }}
      </a-list-item>
    </a-list>
<!--  <div class="result" v-for="item in result">-->
<!--    {{ item }}-->
<!--  </div>-->
    <div class="input-box" id="input" data-wails-no-drag>
      <button class="btn" @click="greet()">Greet</button>
      <button class="btn" @click="fileInfo">Greet</button>
    </div>
    </div>
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
  window.go.main.App.FileInfo(result.value[0]).then((res: any) => {
    console.log(res)
  })
}
</script>

<style>
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
</style>
