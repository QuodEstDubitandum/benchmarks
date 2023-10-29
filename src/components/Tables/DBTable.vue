<script setup lang="ts">
import { ref } from 'vue'
import TestDescriptionVue from '../TestDescription.vue'

const go_select = ref({ time: '', loading: false })
const nodejs_select = ref({ time: '', loading: false })
const python_select = ref({ time: '', loading: false })

const go_insert = ref({ time: '', loading: false })
const nodejs_insert = ref({ time: '', loading: false })
const python_insert = ref({ time: '', loading: false })

const go_update = ref({ time: '', loading: false })
const nodejs_update = ref({ time: '', loading: false })
const python_update = ref({ time: '', loading: false })

const go_delete = ref({ time: '', loading: false })
const nodejs_delete = ref({ time: '', loading: false })
const python_delete = ref({ time: '', loading: false })

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
        python_select.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_select.value.time = time
        go_select.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_select.value.time = time
        nodejs_select.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_select.value.time = time
        python_select.value.loading = false
        break
      case '/db/insert':
        go_insert.value.loading = true
        nodejs_insert.value.loading = true
        python_insert.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_insert.value.time = time
        go_insert.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_insert.value.time = time
        nodejs_insert.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_insert.value.time = time
        python_insert.value.loading = false
        break
      case '/db/update':
        go_update.value.loading = true
        nodejs_update.value.loading = true
        python_update.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_update.value.time = time
        go_update.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_update.value.time = time
        nodejs_update.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_update.value.time = time
        python_update.value.loading = false
        break
      case '/db/delete':
        go_delete.value.loading = true
        nodejs_delete.value.loading = true
        python_delete.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_delete.value.time = time
        go_delete.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_delete.value.time = time
        nodejs_delete.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_delete.value.time = time
        python_delete.value.loading = false
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
    :description="'Perform an INSERT query on a semilarge sqlite file n=10 times'"
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
    :description="'Perform an UPDATE query with a WHERE condition on a semilarge sqlite file n=2 times'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_update.loading"></div>
  <p v-else-if="go_update.time">{{ go_update.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_update.loading"></div>
  <p v-else-if="nodejs_update.time">{{ nodejs_update.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_update.loading"></div>
  <p v-else-if="python_update.time">{{ python_update.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/db/delete'"
    :description="'Perform a DELETE query with a WHERE condition on a semilarge sqlite file n=10 times'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_delete.loading"></div>
  <p v-else-if="go_delete.time">{{ go_delete.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_delete.loading"></div>
  <p v-else-if="nodejs_delete.time">{{ nodejs_delete.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_delete.loading"></div>
  <p v-else-if="python_delete.time">{{ python_delete.time }} ms</p>
  <p v-else></p>
</template>
