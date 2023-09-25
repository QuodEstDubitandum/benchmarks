<script setup lang="ts">
import { ref } from 'vue'
import TestDescriptionVue from '../TestDescription.vue'

const go_simple_loop = ref({ time: '', loading: false })
const nodejs_simple_loop = ref({ time: '', loading: false })
const python_simple_loop = ref({ time: '', loading: false })
const go_simple_loop_parallel = ref({ time: '', loading: false })
const nodejs_simple_loop_parallel = ref({ time: '', loading: false })
const python_simple_loop_parallel = ref({ time: '', loading: false })

// global loading state to disable buttons f.e.
const loading = ref(false)
function setLoading() {
  loading.value = !loading.value
}

// fetch the different backends
async function getData(url: string) {
  // switch case for different tests
  setLoading()
  try {
    await new Promise((res, rej) => {
      setTimeout(res, 1000)
    })
    const res = await fetch(url, { method: 'GET' })
    const json = await res.json()
    console.log(json)
  } catch (err) {
    console.log(err)
  }
  setLoading()
}
</script>

<template>
  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="''"
    :description="'Simple Loop for i=100.000'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_simple_loop.loading"></div>
  <p v-else>{{ go_simple_loop.time }}</p>
  <div class="ping" v-if="nodejs_simple_loop.loading"></div>
  <p v-else>{{ nodejs_simple_loop.time }}</p>
  <div class="ping" v-if="python_simple_loop.loading"></div>
  <p v-else>{{ python_simple_loop.time }}</p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="''"
    :description="'Simple Loop for i=100.000 (multi-threaded)'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_simple_loop_parallel.loading"></div>
  <p v-else>{{ go_simple_loop_parallel.time }}</p>
  <div class="ping" v-if="nodejs_simple_loop_parallel.loading"></div>
  <p v-else>{{ nodejs_simple_loop_parallel.time }}</p>
  <div class="ping" v-if="python_simple_loop_parallel.loading"></div>
  <p v-else>{{ python_simple_loop_parallel.time }}</p>
</template>
