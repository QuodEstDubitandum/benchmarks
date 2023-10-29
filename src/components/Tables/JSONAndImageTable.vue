<script setup lang="ts">
import { ref } from 'vue'
import TestDescriptionVue from '../TestDescription.vue'

const go_serialize = ref({ time: '', loading: false })
const nodejs_serialize = ref({ time: '', loading: false })
const python_serialize = ref({ time: '', loading: false })

const go_deserialize = ref({ time: '', loading: false })
const nodejs_deserialize = ref({ time: '', loading: false })
const python_deserialize = ref({ time: '', loading: false })

const go_decode = ref({ time: '', loading: false })
const nodejs_decode = ref({ time: '', loading: false })
const python_decode = ref({ time: '', loading: false })

const go_encode = ref({ time: '', loading: false })
const nodejs_encode = ref({ time: '', loading: false })
const python_encode = ref({ time: '', loading: false })

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
    const sampleString =
      'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.'
    const deserializeObject = {
      a: sampleString,
      b: sampleString,
      c: sampleString,
      d: sampleString,
      e: sampleString,
      f: sampleString,
      g: sampleString,
      h: sampleString,
      i: sampleString,
      j: sampleString,
      k: sampleString,
      l: sampleString,
      m: sampleString,
      n: sampleString,
      o: sampleString,
      p: sampleString,
      q: sampleString
    }
    switch (url) {
      case '/json/serialize':
        go_serialize.value.loading = true
        nodejs_serialize.value.loading = true
        python_serialize.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_serialize.value.time = time
        go_serialize.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_serialize.value.time = time
        nodejs_serialize.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_serialize.value.time = time
        python_serialize.value.loading = false
        break
      case '/json/deserialize':
        go_deserialize.value.loading = true
        nodejs_deserialize.value.loading = true
        python_deserialize.value.loading = true

        res = await fetch('/go' + url, {
          method: 'POST',
          body: JSON.stringify(deserializeObject),
          headers: { 'Content-Type': 'application/json' }
        })
        time = await res.text()
        go_deserialize.value.time = time
        go_deserialize.value.loading = false

        res = await fetch('/nodejs' + url, {
          method: 'POST',
          body: JSON.stringify(deserializeObject),
          headers: { 'Content-Type': 'application/json' }
        })
        time = await res.text()
        nodejs_deserialize.value.time = time
        nodejs_deserialize.value.loading = false

        res = await fetch('/python' + url + '/', {
          method: 'POST',
          body: JSON.stringify(deserializeObject),
          headers: { 'Content-Type': 'application/json' }
        })
        time = await res.text()
        python_deserialize.value.time = time
        python_deserialize.value.loading = false
        break
      case '/image/decode':
        go_decode.value.loading = true
        nodejs_decode.value.loading = true
        python_decode.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_decode.value.time = time
        go_decode.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_decode.value.time = time
        nodejs_decode.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_decode.value.time = time
        python_decode.value.loading = false
        break
      case '/image/encode':
        go_encode.value.loading = true
        nodejs_encode.value.loading = true
        python_encode.value.loading = true

        res = await fetch('/go' + url, { method: 'GET' })
        time = await res.text()
        go_encode.value.time = time
        go_encode.value.loading = false

        res = await fetch('/nodejs' + url, { method: 'GET' })
        time = await res.text()
        nodejs_encode.value.time = time
        nodejs_encode.value.loading = false

        res = await fetch('/python' + url + '/', { method: 'GET' })
        time = await res.text()
        python_encode.value.time = time
        python_encode.value.loading = false
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
    :url="'/json/serialize'"
    :description="'Serialize a semilarge JSON object n=10.000 times'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_serialize.loading"></div>
  <p v-else-if="go_serialize.time">{{ go_serialize.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_serialize.loading"></div>
  <p v-else-if="nodejs_serialize.time">{{ nodejs_serialize.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_serialize.loading"></div>
  <p v-else-if="python_serialize.time">{{ python_serialize.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/json/deserialize'"
    :description="'Deserialize a semilarge JSON object n=10.000 times'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_deserialize.loading"></div>
  <p v-else-if="go_deserialize.time">{{ go_deserialize.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_deserialize.loading"></div>
  <p v-else-if="nodejs_deserialize.time">{{ nodejs_deserialize.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_deserialize.loading"></div>
  <p v-else-if="python_deserialize.time">{{ python_deserialize.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/image/decode'"
    :description="'Decode a jpg image file'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_decode.loading"></div>
  <p v-else-if="go_decode.time">{{ go_decode.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_decode.loading"></div>
  <p v-else-if="nodejs_decode.time">{{ nodejs_decode.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_decode.loading"></div>
  <p v-else-if="python_decode.time">{{ python_decode.time }} ms</p>
  <p v-else></p>

  <TestDescriptionVue
    :loading="loading"
    :get-data="getData"
    :url="'/image/encode'"
    :description="'Encode a png image file'"
  />
  <div class="vertical-seperator"></div>
  <div class="ping" v-if="go_encode.loading"></div>
  <p v-else-if="go_encode.time">{{ go_encode.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="nodejs_encode.loading"></div>
  <p v-else-if="nodejs_encode.time">{{ nodejs_encode.time }} ms</p>
  <p v-else></p>
  <div class="ping" v-if="python_encode.loading"></div>
  <p v-else-if="python_encode.time">{{ python_encode.time }} ms</p>
  <p v-else></p>
</template>
