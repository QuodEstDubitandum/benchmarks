<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="nav-menu">
    <div class="nav-button"><RouterLink to="/loop">Loops</RouterLink></div>
    <div class="nav-button"><RouterLink to="/serialization">Serialization</RouterLink></div>
    <div class="nav-button"><RouterLink to="/file">File Operations</RouterLink></div>
    <div class="nav-button"><RouterLink to="/db-query">Database Queries</RouterLink></div>
    <div class="nav-button"><RouterLink to="/db-update">Database Updates</RouterLink></div>
  </div>
  <div class="home">
    <RouterLink to="/"><img class="home-icon" src="/home.png" /></RouterLink>
  </div>

  <div class="container">
    <RouterView />
  </div>
</template>

<style scoped>
.container {
  width: 70vw;
  margin: auto;
  display: flex;
  padding: 100px 0 100px 0;
}
.nav-menu {
  display: flex;
  flex-direction: row;
  width: 100vw;
  justify-content: space-between;
  padding: 30px 100px 0px 100px;
}

.nav-button {
  padding: 0 4px 0 4px;
  border-width: 0.5rem;
  border-color: black;
  border-radius: 0.5rem;
  font-size: large;
}

.nav-button > .router-link-active {
  text-decoration: underline; /* This underlines the active link */
  color: #079a1d;
}

a {
  color: white;
  text-decoration: none;
}

.home {
  position: fixed;
  bottom: 30px;
  left: 30px;
  width: 80px;
  height: 80px;
}

.home-icon {
  width: 80px;
  height: 80px;
}
</style>
