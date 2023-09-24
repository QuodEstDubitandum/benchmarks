<template>
  <div class="grid">
    <p>Tests</p>
    <div class="vertical-seperator"></div>
    <p>1</p>
    <p>2</p>
    <p>3</p>
    <div class="horizontal-seperator"></div>
  </div>
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
