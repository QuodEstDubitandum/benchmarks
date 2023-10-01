<script setup lang="ts">
import { ref } from 'vue'
import TestDescriptionVue from '../TestDescription.vue'

const go_write = ref({ time: '', loading: false })
const nodejs_write = ref({ time: '', loading: false })
const python_write = ref({ time: '', loading: false })

const go_write_parallel = ref({ time: '', loading: false })
const nodejs_write_parallel = ref({ time: '', loading: false })
const python_write_parallel = ref({ time: '', loading: false })

const go_quicksort = ref({ time: '', loading: false })
const nodejs_quicksort = ref({ time: '', loading: false })
const python_quicksort = ref({ time: '', loading: false })

const go_two_sum = ref({ time: '', loading: false })
const nodejs_two_sum = ref({ time: '', loading: false })
const python_two_sum = ref({ time: '', loading: false })

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
    let res: Response
    let time: string
    switch (url) {
      case '/files/write':
        go_write.value.loading = true
        nodejs_write.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_write.value.time = time
        go_write.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_write.value.time = time
        nodejs_write.value.loading = false
        break
      case '/files/write-parallel':
        go_write_parallel.value.loading = true
        nodejs_write_parallel.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_write_parallel.value.time = time
        go_write_parallel.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_write_parallel.value.time = time
        nodejs_write_parallel.value.loading = false
        break
      case '/files/quicksort':
        go_quicksort.value.loading = true
        nodejs_quicksort.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_quicksort.value.time = time
        go_quicksort.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_quicksort.value.time = time
        nodejs_quicksort.value.loading = false
        break
      case '/files/two-sum':
        go_two_sum.value.loading = true
        nodejs_two_sum.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_two_sum.value.time = time
        go_two_sum.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_two_sum.value.time = time
        nodejs_two_sum.value.loading = false
        break
    }
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
    :url="'/files/write'"
    :description="'Write n=10.000 lines to NumOfCPUCores text files one after another (single-threaded)'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_write.loading"></div>
  <p v-else-if="go_write.time">{{ go_write.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_write.loading"></div>
  <p v-else-if="nodejs_write.time">{{ nodejs_write.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_write.loading"></div>
  <p v-else-if="python_write.time">{{ python_write.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/files/write-parallel'"
    :description="'Write n=10.000 lines to NumOfCPUCores text files in parallel (multi-threaded)'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_write_parallel.loading"></div>
  <p v-else-if="go_write_parallel.time">{{ go_write_parallel.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_write_parallel.loading"></div>
  <p v-else-if="nodejs_write_parallel.time">{{ nodejs_write_parallel.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_write_parallel.loading"></div>
  <p v-else-if="python_write_parallel.time">{{ python_write_parallel.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/files/quicksort'"
    :description="'Perform Quicksort on a badly sorted array for n=1.000.000'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_quicksort.loading"></div>
  <p v-else-if="go_quicksort.time">{{ go_quicksort.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_quicksort.loading"></div>
  <p v-else-if="nodejs_quicksort.time">{{ nodejs_quicksort.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_quicksort.loading"></div>
  <p v-else-if="python_quicksort.time">{{ python_quicksort.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/files/two-sum'"
    :description="'Given an array of size n=1.000.000: Find the number of pairs in that array whose sum equals n'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_two_sum.loading"></div>
  <p v-else-if="go_two_sum.time">{{ go_two_sum.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_two_sum.loading"></div>
  <p v-else-if="nodejs_two_sum.time">{{ nodejs_two_sum.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_two_sum.loading"></div>
  <p v-else-if="python_two_sum.time">{{ python_two_sum.time }} ms</p>
  <p v-else></p>
</template>
