<script setup lang="ts">
import { ref } from 'vue'

const isOpen = ref(false)

function toggleOpen() {
  isOpen.value = !isOpen.value
}
const props = defineProps({
  title: String,
  description: String
})
</script>

<template>
  <div style="display: flex; align-items: center; margin-bottom: 10px">
    <svg
      @click="toggleOpen"
      xmlns="http://www.w3.org/2000/svg"
      height="1em"
      viewBox="0 0 448 512"
      style="margin-right: 15px; height: 35px; width: 35px; fill: green; cursor: pointer"
      v-if="isOpen"
    >
      <!--! Font Awesome Free 6.4.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
      <path
        d="M201.4 137.4c12.5-12.5 32.8-12.5 45.3 0l160 160c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0L224 205.3 86.6 342.6c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3l160-160z"
      />
    </svg>
    <svg
      @click="toggleOpen"
      xmlns="http://www.w3.org/2000/svg"
      height="1em"
      viewBox="0 0 448 512"
      style="margin-right: 15px; height: 35px; width: 35px; fill: green; cursor: pointer"
      v-else
    >
      <!--! Font Awesome Free 6.4.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
      <path
        d="M201.4 342.6c12.5 12.5 32.8 12.5 45.3 0l160-160c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L224 274.7 86.6 137.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l160 160z"
      />
    </svg>
    <h2>{{ props.title }}</h2>
  </div>
  <p v-if="isOpen" style="margin-bottom: 10px">{{ props.description }}</p>
</template>
