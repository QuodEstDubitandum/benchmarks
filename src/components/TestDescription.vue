<script setup lang="ts">
const props = defineProps({
  loading: Boolean,
  getData: Function,
  url: String,
  description: String
})
</script>

<template>
  <div class="col">
    <p style="font-size: small">{{ props.description }}</p>
    <div style="margin-bottom: 27px"></div>
    <i
      class="fas fa-circle-play fa-xl"
      style="color: green; cursor: pointer"
      v-if="!props.loading"
      @click="props.getData!(props.url)"
    ></i>
    <i class="fas fa-circle-play fa-xl" style="color: grey" v-else></i>
  </div>
</template>
