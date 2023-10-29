<script setup lang="ts">
import { ref } from 'vue'
import TestDescriptionVue from '../TestDescription.vue'

const go_prime = ref({ time: '', loading: false })
const nodejs_prime = ref({ time: '', loading: false })
const python_prime = ref({ time: '', loading: false })

const go_fibonacci = ref({ time: '', loading: false })
const nodejs_fibonacci = ref({ time: '', loading: false })
const python_fibonacci = ref({ time: '', loading: false })

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
      case '/algorithms/prime':
        go_prime.value.loading = true
        nodejs_prime.value.loading = true
        python_prime.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_prime.value.time = time
        go_prime.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_prime.value.time = time
        nodejs_prime.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_prime.value.time = time
        python_prime.value.loading = false
        break
      case '/algorithms/fast-fibonacci':
        go_fibonacci.value.loading = true
        nodejs_fibonacci.value.loading = true
        python_fibonacci.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_fibonacci.value.time = time
        go_fibonacci.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_fibonacci.value.time = time
        nodejs_fibonacci.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_fibonacci.value.time = time
        python_fibonacci.value.loading = false
        break
      case '/algorithms/quicksort':
        go_quicksort.value.loading = true
        nodejs_quicksort.value.loading = true
        python_quicksort.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_quicksort.value.time = time
        go_quicksort.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_quicksort.value.time = time
        nodejs_quicksort.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_quicksort.value.time = time
        python_quicksort.value.loading = false
        break
      case '/algorithms/two-sum':
        go_two_sum.value.loading = true
        nodejs_two_sum.value.loading = true
        python_two_sum.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_two_sum.value.time = time
        go_two_sum.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_two_sum.value.time = time
        nodejs_two_sum.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_two_sum.value.time = time
        python_two_sum.value.loading = false
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
    :url="'/algorithms/prime'"
    :description="'Calculate all prime numbers up to n=10.000.000'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_prime.loading"></div>
  <p v-else-if="go_prime.time">{{ go_prime.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_prime.loading"></div>
  <p v-else-if="nodejs_prime.time">{{ nodejs_prime.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_prime.loading"></div>
  <p v-else-if="python_prime.time">{{ python_prime.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/algorithms/fast-fibonacci'"
    :description="'Compute the fibonacci number for n=1.000.000'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_fibonacci.loading"></div>
  <p v-else-if="go_fibonacci.time">{{ go_fibonacci.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_fibonacci.loading"></div>
  <p v-else-if="nodejs_fibonacci.time">{{ nodejs_fibonacci.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_fibonacci.loading"></div>
  <p v-else-if="python_fibonacci.time">{{ python_fibonacci.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/algorithms/quicksort'"
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
    :url="'/algorithms/two-sum'"
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
