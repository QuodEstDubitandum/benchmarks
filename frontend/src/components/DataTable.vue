<script setup lang="ts">
import LoopTable from '../components/Tables/LoopTable.vue'
</script>

<template>
  <main style="width: 100%; height: fit-content; position: relative">
    <div class="grid">
      <p>Tests</p>
      <div class="vertical-seperator" style="top: 0"></div>
      <p>Go</p>
      <p>NodeJS</p>
      <p>Python</p>
      <div class="horizontal-seperator"></div>
      <div v-if="$route.path === '/loop'"><LoopTable /></div>
      <div v-else-if="$route.path === '/serialization'"></div>
      <div v-else-if="$route.path === '/file'"></div>
      <div v-else-if="$route.path === '/db-query'"></div>
      <div v-else></div>
    </div>
  </main>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(1, 1fr) 1px repeat(3, 1fr);
  color: white;
  justify-items: center;
  align-content: center;
  gap: 10px;
}
</style>
