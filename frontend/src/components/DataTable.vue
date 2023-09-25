<script setup lang="ts">
import LoopTable from '../components/Tables/LoopTable.vue'
</script>

<template>
  <main style="width: 100%; height: fit-content; position: relative">
    <div class="grid" style="row-gap: 20px">
      <p>Tests</p>
      <div class="vertical-seperator" style="top: 20px"></div>
      <p>Go</p>
      <p>NodeJS</p>
      <p>Python</p>
      <div class="horizontal-seperator"></div>
    </div>
    <div style="width: 100%; height: 40px"></div>
    <div class="grid">
      <LoopTable v-if="$route.path === '/loop'" />
      <LoopTable v-if="$route.path === '/serialization'" />
      <LoopTable v-if="$route.path === '/file'" />
      <LoopTable v-if="$route.path === '/db-query'" />
      <LoopTable v-if="$route.path === '/db-update'" />
    </div>
  </main>
</template>

<style scoped></style>
