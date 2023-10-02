<script setup lang="ts">
import { ref } from 'vue'
import TestDescriptionVue from '../TestDescription.vue'

const go_write = ref({ time: '', loading: false })
const nodejs_write = ref({ time: '', loading: false })
const python_write = ref({ time: '', loading: false })

const go_write_parallel = ref({ time: '', loading: false })
const nodejs_write_parallel = ref({ time: '', loading: false })
const python_write_parallel = ref({ time: '', loading: false })

const go_read = ref({ time: '', loading: false })
const nodejs_read = ref({ time: '', loading: false })
const python_read = ref({ time: '', loading: false })

const go_read_parallel = ref({ time: '', loading: false })
const nodejs_read_parallel = ref({ time: '', loading: false })
const python_read_parallel = ref({ time: '', loading: false })

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
      case '/files/read':
        go_read.value.loading = true
        nodejs_read.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_read.value.time = time
        go_read.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_read.value.time = time
        nodejs_read.value.loading = false
        break
      case '/files/read-parallel':
        go_read_parallel.value.loading = true
        nodejs_read_parallel.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_read_parallel.value.time = time
        go_read_parallel.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_read_parallel.value.time = time
        nodejs_read_parallel.value.loading = false
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
    :description="'Write n=10.000 lines to NumOfCPUCores text files one after another'"
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
    :description="'Write n=10.000 lines to NumOfCPUCores text files in parallel'"
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
    :url="'/files/read'"
    :description="'Read NumOfCPUCores text files consisting of n=1.000.000 lines each synchronously'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_read.loading"></div>
  <p v-else-if="go_read.time">{{ go_read.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_read.loading"></div>
  <p v-else-if="nodejs_read.time">{{ nodejs_read.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_read.loading"></div>
  <p v-else-if="python_read.time">{{ python_read.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/files/read-parallel'"
    :description="'Read NumOfCPUCores text files consisting of n=1.000.000 lines each in parallel'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_read_parallel.loading"></div>
  <p v-else-if="go_read_parallel.time">{{ go_read_parallel.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_read_parallel.loading"></div>
  <p v-else-if="nodejs_read_parallel.time">{{ nodejs_read_parallel.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_read_parallel.loading"></div>
  <p v-else-if="python_read_parallel.time">{{ python_read_parallel.time }} ms</p>
  <p v-else></p>
</template>
