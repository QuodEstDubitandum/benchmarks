<script setup lang="ts">
import { ref } from 'vue'
import TestDescriptionVue from '../TestDescription.vue'

const go_select = ref({ time: '', loading: false })
const nodejs_select = ref({ time: '', loading: false })
const python_select = ref({ time: '', loading: false })

const go_insert = ref({ time: '', loading: false })
const nodejs_insert = ref({ time: '', loading: false })
const python_insert = ref({ time: '', loading: false })

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
      case '/db/select':
        go_select.value.loading = true
        nodejs_select.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_select.value.time = time
        go_select.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_select.value.time = time
        nodejs_select.value.loading = false
        break
      case '/db/insert':
        go_insert.value.loading = true
        nodejs_insert.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_insert.value.time = time
        go_insert.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_insert.value.time = time
        nodejs_insert.value.loading = false
        break
      case '/db/update':
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
      case '/db/delete':
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
    :url="'/db/select'"
    :description="'Perform a SELECT query with a WHERE condition on a semilarge sqlite file n=10 times'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_select.loading"></div>
  <p v-else-if="go_select.time">{{ go_select.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_select.loading"></div>
  <p v-else-if="nodejs_select.time">{{ nodejs_select.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_select.loading"></div>
  <p v-else-if="python_select.time">{{ python_select.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/db/insert'"
    :description="'Perform an INSERT query on a semilarge sqlite file n=1.000 times'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_insert.loading"></div>
  <p v-else-if="go_insert.time">{{ go_insert.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_insert.loading"></div>
  <p v-else-if="nodejs_insert.time">{{ nodejs_insert.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_insert.loading"></div>
  <p v-else-if="python_insert.time">{{ python_insert.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/db/update'"
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
    :url="'/db/delete'"
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
