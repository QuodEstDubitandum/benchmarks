<script setup lang="ts">
import AlgorithmTable from '../components/Tables/AlgorithmTable.vue'
import FileTable from '../components/Tables/FileTable.vue'
import DBTable from '../components/Tables/DBTable.vue'
import JSONAndImageTable from '../components/Tables/JSONAndImageTable.vue'
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
      <AlgorithmTable v-if="$route.path === '/algorithms'" />
      <FileTable v-if="$route.path === '/file'" />
      <DBTable v-if="$route.path === '/db'" />
      <JSONAndImageTable v-if="$route.path === '/json-image'" />
    </div>
  </main>
</template>

<style scoped></style>
